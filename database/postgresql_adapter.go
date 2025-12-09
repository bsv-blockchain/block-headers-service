package database

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// use blank import to use file source driver with the migrate package.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/block-headers-service/config"
	"github.com/bsv-blockchain/block-headers-service/database/sql"
	"github.com/bsv-blockchain/block-headers-service/internal/chaincfg/chainhash"
	"github.com/bsv-blockchain/block-headers-service/repository/dto"
)

type postgreSQLAdapter struct {
	db *sqlx.DB
}

const (
	postgresDriverName = "postgres"
	postgresBatchSize  = 500_000
)

func (a *postgreSQLAdapter) connect(cfg *config.DbConfig) error {
	dbCfg := cfg.Postgres
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName, dbCfg.Sslmode)

	db, err := sqlx.Open(postgresDriverName, dsn)
	if err != nil {
		return err
	}

	a.db = db
	return nil
}

func (a *postgreSQLAdapter) doMigrations(cfg *config.DbConfig) error {
	driver, err := postgres.WithInstance(a.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	sourceURL := fmt.Sprintf("file://%s", cfg.SchemaPath)

	m, err := migrate.NewWithDatabaseInstance(sourceURL, postgresDriverName, driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (a *postgreSQLAdapter) getDBx() *sqlx.DB {
	if a.db == nil {
		panic("connection to the database has not been established")
	}
	return a.db
}

func (a *postgreSQLAdapter) importHeaders(inputFile *os.File, _ *zerolog.Logger) (affectedRows int, err error) {
	// prepare db for bulk insterts
	restoreIndexes, err := a.dropTableIndexes(sql.HeadersTableName)
	if err != nil {
		return affectedRows, err
	}
	defer func() {
		if rErr := restoreIndexes(); rErr != nil {
			err = wrapIfNeeded(err, rErr, "Resoring indexes failed")
		}
	}()

	if _, err = inputFile.Seek(0, 0); err != nil {
		return affectedRows, err
	}

	reader := csv.NewReader(inputFile)
	_, err = reader.Read() // Skipping the column headers line
	if err != nil {
		return affectedRows, err
	}

	// insert headers
	previousBlockHash := chainhash.Hash{}.String()
	var cumulatedChainWork string
	rowIndex := 0
	guard := 0

	for {
		rowIndex, previousBlockHash, cumulatedChainWork, err = a.copyHeaders(reader, postgresBatchSize, previousBlockHash, cumulatedChainWork, rowIndex)
		if err != nil {
			affectedRows = rowIndex
			return affectedRows, err
		}

		if guard == rowIndex {
			break
		}

		guard = rowIndex
		affectedRows = rowIndex
	}

	return affectedRows, err
}

// dropTableIndexes removes indexes from a table. Returns the index restore function if successful.
func (a *postgreSQLAdapter) dropTableIndexes(table string) (func() error, error) {
	q := fmt.Sprintf("SELECT indexname, indexdef FROM pg_indexes WHERE tablename ='%s' AND indexname != '%s_pkey' AND indexdef IS NOT NULL;", table, table)
	return dropIndexes(a.db, &q)
}

func (a *postgreSQLAdapter) copyHeaders(reader *csv.Reader, batchSize int, previousBlockHash, cumulatedLastBlockChainWork string, rowIndex int) (lastRowIndex int, lastBlockHash, cumulatedChainWork string, err error) {
	lastRowIndex = rowIndex
	lastBlockHash = previousBlockHash
	copyQuery := pq.CopyIn(
		sql.HeadersTableName,
		/* columns */ "height", "hash", "version", "merkleroot", "timestamp", "bits", "nonce", "header_state", "chainwork", "cumulated_work", "previous_block",
	)

	dbTx, err := a.db.Begin()
	if err != nil {
		return lastRowIndex, lastBlockHash, cumulatedChainWork, err
	}
	defer dbTx.Rollback() //nolint

	stmt, err := dbTx.Prepare(copyQuery)
	if err != nil {
		return lastRowIndex, lastBlockHash, cumulatedChainWork, err
	}

	cumulatedChainWork = cumulatedLastBlockChainWork
	for i := 0; i < batchSize; i++ {
		record, readErr := reader.Read()
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				break
			}

			err = fmt.Errorf("error reading record: %v", readErr)
			_ = stmt.Close() //nolint
			return lastRowIndex, lastBlockHash, cumulatedChainWork, err
		}

		if len(record) == 0 {
			break
		}
		var b *dto.DbBlockHeader
		b, err = prepareRecord(record, lastBlockHash, cumulatedChainWork, lastRowIndex)
		if err != nil {
			return lastRowIndex, lastBlockHash, cumulatedChainWork, err
		}

		_, execErr := stmt.Exec(
			b.Height,
			b.Hash,
			b.Version,
			b.MerkleRoot,
			b.Timestamp,
			b.Bits,
			b.Nonce,
			b.State,
			b.Chainwork,
			b.CumulatedWork,
			b.PreviousBlock)

		if execErr != nil {
			err = fmt.Errorf("error preparing copy statement after %d row: %v", lastRowIndex, execErr)
			return lastRowIndex, lastBlockHash, cumulatedChainWork, err
		}

		cumulatedChainWork = b.CumulatedWork
		lastBlockHash = b.Hash
		lastRowIndex++
	}

	_, err = stmt.Exec()
	if err != nil {
		if closeErr := stmt.Close(); closeErr != nil {
			err = fmt.Errorf("execution err: %w. Smt close err: %w", err, closeErr)
		}
		return lastRowIndex, lastBlockHash, cumulatedChainWork, err
	}

	err = stmt.Close()
	if err != nil {
		return lastRowIndex, lastBlockHash, cumulatedChainWork, err
	}

	err = dbTx.Commit()
	return lastRowIndex, lastBlockHash, cumulatedChainWork, err
}
