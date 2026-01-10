package dto

import (
	"database/sql"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bsv-blockchain/block-headers-service/domains"
)

func TestConvertToBlockHeader(t *testing.T) {
	tests := []struct {
		name     string
		input    []*DbBlockHeader
		expected int
	}{
		{
			name:     "nil input returns nil",
			input:    nil,
			expected: 0,
		},
		{
			name:     "empty slice returns empty slice",
			input:    []*DbBlockHeader{},
			expected: 0,
		},
		{
			name: "single element conversion",
			input: []*DbBlockHeader{
				{
					Height:        100,
					Hash:          "0000000000000000000000000000000000000000000000000000000000000001",
					Version:       1,
					MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000002",
					Timestamp:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					Bits:          486604799,
					Nonce:         123456,
					State:         "LONGEST_CHAIN",
					Chainwork:     "100",
					CumulatedWork: "1000",
					PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000000",
				},
			},
			expected: 1,
		},
		{
			name: "multiple elements conversion",
			input: []*DbBlockHeader{
				{
					Height:        100,
					Hash:          "0000000000000000000000000000000000000000000000000000000000000001",
					Version:       1,
					MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000002",
					Timestamp:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					Bits:          486604799,
					Nonce:         123456,
					State:         "LONGEST_CHAIN",
					Chainwork:     "100",
					CumulatedWork: "1000",
					PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000000",
				},
				{
					Height:        101,
					Hash:          "0000000000000000000000000000000000000000000000000000000000000003",
					Version:       1,
					MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000004",
					Timestamp:     time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
					Bits:          486604799,
					Nonce:         789012,
					State:         "LONGEST_CHAIN",
					Chainwork:     "100",
					CumulatedWork: "1100",
					PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000001",
				},
				{
					Height:        102,
					Hash:          "0000000000000000000000000000000000000000000000000000000000000005",
					Version:       1,
					MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000006",
					Timestamp:     time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
					Bits:          486604799,
					Nonce:         345678,
					State:         "STALE",
					Chainwork:     "100",
					CumulatedWork: "1200",
					PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000003",
				},
			},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertToBlockHeader(tt.input)

			if tt.input == nil {
				assert.Nil(t, result)
				return
			}

			require.Len(t, result, tt.expected)

			for i, header := range result {
				assert.Equal(t, tt.input[i].Height, header.Height)
				assert.Equal(t, tt.input[i].Hash, header.Hash.String())
				assert.Equal(t, tt.input[i].Version, header.Version)
				assert.Equal(t, tt.input[i].MerkleRoot, header.MerkleRoot.String())
				assert.Equal(t, tt.input[i].Bits, header.Bits)
				assert.Equal(t, tt.input[i].Nonce, header.Nonce)
				assert.Equal(t, domains.HeaderState(tt.input[i].State), header.State)
			}
		})
	}
}

func TestConvertToBlockHeader_FieldValues(t *testing.T) {
	input := []*DbBlockHeader{
		{
			Height:        12345,
			Hash:          "00000000000000000008a928c3f5b98f3b53bbbdc6f4e0b093e9d2f0a0989c4c",
			Version:       536870912,
			MerkleRoot:    "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b",
			Timestamp:     time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC),
			Bits:          386089497,
			Nonce:         2083236893,
			State:         "LONGEST_CHAIN",
			Chainwork:     "12345678901234567890",
			CumulatedWork: "98765432109876543210",
			PreviousBlock: "00000000000000000007316856900e76b4f7a9139cfbfba89842c8d196cd5f91",
		},
	}

	result := ConvertToBlockHeader(input)

	require.Len(t, result, 1)
	header := result[0]

	assert.Equal(t, int32(12345), header.Height)
	assert.Equal(t, int32(536870912), header.Version)
	assert.Equal(t, uint32(386089497), header.Bits)
	assert.Equal(t, uint32(2083236893), header.Nonce)
	assert.Equal(t, domains.LongestChain, header.State)

	expectedChainwork, _ := new(big.Int).SetString("12345678901234567890", 10)
	expectedCumulatedWork, _ := new(big.Int).SetString("98765432109876543210", 10)
	assert.Equal(t, expectedChainwork, header.Chainwork)
	assert.Equal(t, expectedCumulatedWork, header.CumulatedWork)
}

func TestConvertToMerkleRootsConfirmations(t *testing.T) {
	tests := []struct {
		name                 string
		input                []*DbMerkleRootConfirmation
		maxBlockHeightExcess int
		expected             int
	}{
		{
			name:                 "empty slice returns empty slice",
			input:                []*DbMerkleRootConfirmation{},
			maxBlockHeightExcess: 6,
			expected:             0,
		},
		{
			name: "single confirmed element",
			input: []*DbMerkleRootConfirmation{
				{
					MerkleRoot:  "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b",
					BlockHeight: 100,
					Hash:        sql.NullString{String: "00000000000000000001", Valid: true},
					TipHeight:   100,
				},
			},
			maxBlockHeightExcess: 6,
			expected:             1,
		},
		{
			name: "multiple elements with different states",
			input: []*DbMerkleRootConfirmation{
				{
					MerkleRoot:  "merkle1",
					BlockHeight: 100,
					Hash:        sql.NullString{String: "hash1", Valid: true},
					TipHeight:   105,
				},
				{
					MerkleRoot:  "merkle2",
					BlockHeight: 110,
					Hash:        sql.NullString{Valid: false},
					TipHeight:   105,
				},
				{
					MerkleRoot:  "merkle3",
					BlockHeight: 200,
					Hash:        sql.NullString{Valid: false},
					TipHeight:   105,
				},
			},
			maxBlockHeightExcess: 6,
			expected:             3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertToMerkleRootsConfirmations(tt.input, tt.maxBlockHeightExcess)

			require.Len(t, result, tt.expected)

			for i, confm := range result {
				assert.Equal(t, tt.input[i].MerkleRoot, confm.MerkleRoot)
				assert.Equal(t, tt.input[i].BlockHeight, confm.BlockHeight)
			}
		})
	}
}

func TestConvertToMerkleRootsConfirmations_States(t *testing.T) {
	tests := []struct {
		name                 string
		input                *DbMerkleRootConfirmation
		maxBlockHeightExcess int
		expectedState        domains.MerkleRootConfirmationState
	}{
		{
			name: "confirmed when hash is valid",
			input: &DbMerkleRootConfirmation{
				MerkleRoot:  "merkle1",
				BlockHeight: 100,
				Hash:        sql.NullString{String: "valid_hash", Valid: true},
				TipHeight:   105,
			},
			maxBlockHeightExcess: 6,
			expectedState:        domains.Confirmed,
		},
		{
			name: "unable to verify when block height exceeds tip but within excess",
			input: &DbMerkleRootConfirmation{
				MerkleRoot:  "merkle2",
				BlockHeight: 108,
				Hash:        sql.NullString{Valid: false},
				TipHeight:   105,
			},
			maxBlockHeightExcess: 6,
			expectedState:        domains.UnableToVerify,
		},
		{
			name: "invalid when block height far exceeds tip",
			input: &DbMerkleRootConfirmation{
				MerkleRoot:  "merkle3",
				BlockHeight: 200,
				Hash:        sql.NullString{Valid: false},
				TipHeight:   105,
			},
			maxBlockHeightExcess: 6,
			expectedState:        domains.Invalid,
		},
		{
			name: "invalid when hash not valid and block height at or below tip",
			input: &DbMerkleRootConfirmation{
				MerkleRoot:  "merkle4",
				BlockHeight: 100,
				Hash:        sql.NullString{Valid: false},
				TipHeight:   105,
			},
			maxBlockHeightExcess: 6,
			expectedState:        domains.Invalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertToMerkleRootsConfirmations(
				[]*DbMerkleRootConfirmation{tt.input},
				tt.maxBlockHeightExcess,
			)

			require.Len(t, result, 1)
			assert.Equal(t, tt.expectedState, result[0].Confirmation)
		})
	}
}

func TestDbBlockHeader_ToBlockHeader(t *testing.T) {
	dbh := &DbBlockHeader{
		Height:        500000,
		Hash:          "0000000000000000000000000000000000000000000000000000000000000001",
		Version:       536870912,
		MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000002",
		Timestamp:     time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		Bits:          386089497,
		Nonce:         2083236893,
		State:         "LONGEST_CHAIN",
		Chainwork:     "1234567890",
		CumulatedWork: "9876543210",
		PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000000",
	}

	result := dbh.ToBlockHeader()

	assert.NotNil(t, result)
	assert.Equal(t, dbh.Height, result.Height)
	assert.Equal(t, dbh.Version, result.Version)
	assert.Equal(t, dbh.Bits, result.Bits)
	assert.Equal(t, dbh.Nonce, result.Nonce)
	assert.Equal(t, domains.LongestChain, result.State)
}

func TestDbBlockHeader_ToBlockHeader_EmptyCumulatedWork(t *testing.T) {
	dbh := &DbBlockHeader{
		Height:        100,
		Hash:          "0000000000000000000000000000000000000000000000000000000000000001",
		MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000002",
		PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000000",
		CumulatedWork: "",
		Chainwork:     "100",
	}

	result := dbh.ToBlockHeader()

	assert.NotNil(t, result)
	assert.Equal(t, big.NewInt(0), result.CumulatedWork)
}

func TestDbBlockHeader_ToBlockHeader_InvalidChainwork(t *testing.T) {
	dbh := &DbBlockHeader{
		Height:        100,
		Hash:          "0000000000000000000000000000000000000000000000000000000000000001",
		MerkleRoot:    "0000000000000000000000000000000000000000000000000000000000000002",
		PreviousBlock: "0000000000000000000000000000000000000000000000000000000000000000",
		CumulatedWork: "invalid",
		Chainwork:     "also_invalid",
	}

	result := dbh.ToBlockHeader()

	assert.NotNil(t, result)
	assert.Equal(t, big.NewInt(0), result.CumulatedWork)
	assert.Equal(t, big.NewInt(0), result.Chainwork)
}

func TestToDbBlockHeader(t *testing.T) {
	chainwork, _ := new(big.Int).SetString("1234567890", 10)
	cumulatedWork, _ := new(big.Int).SetString("9876543210", 10)

	bh := domains.BlockHeader{
		Height:        100,
		Version:       1,
		Timestamp:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Bits:          486604799,
		Nonce:         123456,
		State:         domains.LongestChain,
		Chainwork:     chainwork,
		CumulatedWork: cumulatedWork,
	}

	result := ToDbBlockHeader(bh)

	assert.Equal(t, bh.Height, result.Height)
	assert.Equal(t, bh.Version, result.Version)
	assert.Equal(t, bh.Timestamp, result.Timestamp)
	assert.Equal(t, bh.Bits, result.Bits)
	assert.Equal(t, bh.Nonce, result.Nonce)
	assert.Equal(t, "LONGEST_CHAIN", result.State)
	assert.Equal(t, "1234567890", result.Chainwork)
	assert.Equal(t, "9876543210", result.CumulatedWork)
}

func TestDbMerkleRootConfirmation_ToMerkleRootConfirmation(t *testing.T) {
	dbConfm := &DbMerkleRootConfirmation{
		MerkleRoot:  "test_merkle_root",
		BlockHeight: 12345,
		Hash:        sql.NullString{String: "test_hash", Valid: true},
		TipHeight:   12350,
	}

	result := dbConfm.ToMerkleRootConfirmation(6)

	assert.Equal(t, "test_merkle_root", result.MerkleRoot)
	assert.Equal(t, int32(12345), result.BlockHeight)
	assert.Equal(t, "test_hash", result.Hash)
	assert.Equal(t, domains.Confirmed, result.Confirmation)
}
