package config

import (
	"testing"

	"github.com/bsv-blockchain/block-headers-service/internal/tests/assert"
)

func TestP2PConfig_GetListenPort(t *testing.T) {
	tests := map[string]struct {
		cfg  *P2PConfig
		want string
	}{
		"unset uses the network default": {cfg: &P2PConfig{}, want: "8333"},
		"configured port":                {cfg: &P2PConfig{ListenPort: 18333}, want: "18333"},
		"port above 32767":               {cfg: &P2PConfig{ListenPort: 48333}, want: "48333"},
		"nil config uses the default":    {cfg: nil, want: "8333"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.cfg.GetListenPort("8333"), tt.want)
		})
	}
}

func TestP2PConfig_ValidateListenPort(t *testing.T) {
	tests := map[string]struct {
		port    int
		wantErr bool
	}{
		"zero (network default)": {port: 0},
		"typical":                {port: 18333},
		"highest":                {port: 65535},
		"negative":               {port: -1, wantErr: true},
		"too large":              {port: 65536, wantErr: true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := (&P2PConfig{ListenPort: tt.port}).Validate()
			if tt.wantErr && err == nil {
				t.Fatalf("expected an error for port %d", tt.port)
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error for port %d: %v", tt.port, err)
			}
		})
	}
}

func TestDefaultConfig_ListenPortIsNetworkDefault(t *testing.T) {
	assert.Equal(t, getP2PDefaults().ListenPort, 0)
}
