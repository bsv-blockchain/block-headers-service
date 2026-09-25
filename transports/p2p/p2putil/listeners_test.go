package p2putil

import (
	"net"
	"strconv"
	"testing"

	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/block-headers-service/internal/tests/assert"
)

func TestInitListeners(t *testing.T) {
	// given
	log := zerolog.Nop()

	// when
	listeners, err := InitListeners(&log, "8333")

	// then
	assert.NoError(t, err)
	assert.Equal(t, len(listeners), 1)
	assert.Equal(t, listeners[0].Addr().Network(), "tcp")
	assert.Equal(t, listeners[0].Addr().String(), "[::]:8333")
}

// A configured port is used instead of the network default (#324).
func TestInitListeners_CustomPort(t *testing.T) {
	// given
	log := zerolog.Nop()
	probe, err := net.Listen("tcp", "[::]:0")
	assert.NoError(t, err)
	port := strconv.Itoa(probe.Addr().(*net.TCPAddr).Port)
	assert.NoError(t, probe.Close())

	// when
	listeners, err := InitListeners(&log, port)

	// then
	assert.NoError(t, err)
	assert.Equal(t, len(listeners), 1)
	assert.Equal(t, listeners[0].Addr().String(), "[::]:"+port)
	for _, l := range listeners {
		assert.NoError(t, l.Close())
	}
}
