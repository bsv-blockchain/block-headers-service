package p2putil

import (
	"net"
	"strconv"
	"testing"

	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/block-headers-service/internal/tests/assert"
)

// freePort returns a TCP port that was free a moment ago.
func freePort(t *testing.T) string {
	t.Helper()
	probe, err := net.Listen("tcp", "[::]:0")
	assert.NoError(t, err)
	port := strconv.Itoa(probe.Addr().(*net.TCPAddr).Port)
	assert.NoError(t, probe.Close())
	return port
}

// initAndClose starts listeners on port, checks each is TCP on that port, and
// closes them. Dual-stack hosts may give one listener or two.
func initAndClose(t *testing.T, port string) {
	t.Helper()
	log := zerolog.Nop()

	listeners, err := InitListeners(&log, port)

	assert.NoError(t, err)
	assert.Equal(t, len(listeners) > 0, true)
	for _, l := range listeners {
		assert.Equal(t, l.Addr().Network(), "tcp")
		assert.Equal(t, strconv.Itoa(l.Addr().(*net.TCPAddr).Port), port)
		assert.NoError(t, l.Close())
	}
}

// InitListeners listens on the port it is given, not a network default (#324).
func TestInitListeners(t *testing.T) {
	initAndClose(t, freePort(t))
}
