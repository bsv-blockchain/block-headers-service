package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	peerpkg "github.com/bsv-blockchain/block-headers-service/transports/p2p/peer"
)

func TestNewNetworkService(t *testing.T) {
	peers := make(map[*peerpkg.Peer]*peerpkg.SyncState)

	service := NewNetworkService(peers)

	assert.NotNil(t, service)
	assert.Equal(t, peers, service.peers)
}

func TestNetworkService_GetPeersCount_EmptyPeers(t *testing.T) {
	peers := make(map[*peerpkg.Peer]*peerpkg.SyncState)
	service := NewNetworkService(peers)

	count := service.GetPeersCount()

	assert.Equal(t, 0, count)
}

func TestNetworkService_GetPeersCount_NilPeers(t *testing.T) {
	service := NewNetworkService(nil)

	count := service.GetPeersCount()

	assert.Equal(t, 0, count)
}

func TestNetworkService_GetPeers_EmptyPeers(t *testing.T) {
	peers := make(map[*peerpkg.Peer]*peerpkg.SyncState)
	service := NewNetworkService(peers)

	result := service.GetPeers()

	assert.NotNil(t, result)
	assert.Empty(t, result)
}

func TestNetworkService_GetPeers_NilPeers(t *testing.T) {
	service := NewNetworkService(nil)

	result := service.GetPeers()

	assert.NotNil(t, result)
	assert.Empty(t, result)
}
