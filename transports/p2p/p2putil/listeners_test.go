package p2putil

import (
	"testing"

	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/block-headers-service/internal/tests/assert"
)

func TestInitListeners(t *testing.T) {
	// given
	log := zerolog.Nop()

	// when
	listeners, err := InitListeners(&log)

	// then
	assert.NoError(t, err)
	assert.Equal(t, len(listeners), 1)
	assert.Equal(t, listeners[0].Addr().Network(), "tcp")
	assert.Equal(t, listeners[0].Addr().String(), "[::]:8333")
}
