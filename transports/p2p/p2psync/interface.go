// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package p2psync

import (
	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/block-headers-service/internal/chaincfg"
	"github.com/bsv-blockchain/block-headers-service/internal/chaincfg/chainhash"
	"github.com/bsv-blockchain/block-headers-service/internal/wire"
	"github.com/bsv-blockchain/block-headers-service/service"
	"github.com/bsv-blockchain/block-headers-service/transports/p2p/peer"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	BanPeer(sp *peer.Peer)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	Logger       *zerolog.Logger
	PeerNotifier PeerNotifier
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	MinSyncPeerNetworkSpeed   uint64
	BlocksForForkConfirmation int

	Services    *service.Services
	Checkpoints []chaincfg.Checkpoint
}
