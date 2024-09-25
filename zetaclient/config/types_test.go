package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zeta-chain/node/pkg/chains"
	"github.com/zeta-chain/node/zetaclient/config"
)

func Test_GetRelayerKeyPath(t *testing.T) {
	// create config
	cfg := config.New(false)

	// should return default relayer key path
	require.Equal(t, config.DefaultRelayerKeyPath, cfg.GetRelayerKeyPath())
}

func Test_GetEVMConfig(t *testing.T) {
	chain := chains.Sepolia
	chainID := chains.Sepolia.ChainId

	t.Run("should find non-empty evm config", func(t *testing.T) {
		// create config with defaults
		cfg := config.New(true)

		// set valid evm endpoint
		cfg.EVMChainConfigs[chainID] = config.EVMConfig{
			Chain:    chain,
			Endpoint: "localhost",
		}

		// should return non-empty evm config
		evmCfg, found := cfg.GetEVMConfig(chainID)
		require.True(t, found)
		require.False(t, evmCfg.Empty())
	})

	t.Run("should not find evm config if endpoint is empty", func(t *testing.T) {
		// create config with defaults
		cfg := config.New(true)

		// should not find evm config because endpoint is empty
		_, found := cfg.GetEVMConfig(chainID)
		require.False(t, found)
	})

	t.Run("should not find evm config if chain is empty", func(t *testing.T) {
		// create config with defaults
		cfg := config.New(true)

		// set empty chain
		cfg.EVMChainConfigs[chainID] = config.EVMConfig{
			Chain:    chains.Chain{},
			Endpoint: "localhost",
		}

		// should not find evm config because chain is empty
		_, found := cfg.GetEVMConfig(chainID)
		require.False(t, found)
	})
}

func Test_GetBTCConfig(t *testing.T) {
	tests := []struct {
		name    string
		chainID int64
		oldCfg  config.BTCConfig
		btcCfg  *config.BTCConfig
		want    bool
	}{
		{
			name:    "should find non-empty btc config",
			chainID: chains.BitcoinRegtest.ChainId,
			btcCfg: &config.BTCConfig{
				RPCHost: "localhost",
			},
			want: true,
		},
		{
			name:    "should fallback to old 'BitcoinConfig' if new config is not set",
			chainID: chains.BitcoinRegtest.ChainId,
			oldCfg: config.BTCConfig{
				RPCHost: "old_host",
			},
			btcCfg: nil, // new config is not set
			want:   true,
		},
		{
			name:    "should fallback to old config but still can't find btc config as it's empty",
			chainID: chains.BitcoinRegtest.ChainId,
			oldCfg: config.BTCConfig{
				RPCUsername: "user",
				RPCPassword: "pass",
				RPCHost:     "", // empty config
				RPCParams:   "regtest",
			},
			btcCfg: &config.BTCConfig{
				RPCUsername: "user",
				RPCPassword: "pass",
				RPCHost:     "", // empty config
				RPCParams:   "regtest",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create config with defaults
			cfg := config.New(true)

			// set both new and old btc config
			cfg.BitcoinConfig = tt.oldCfg
			if tt.btcCfg != nil {
				cfg.BTCChainConfigs[tt.chainID] = *tt.btcCfg
			}

			// should return btc config
			btcCfg, found := cfg.GetBTCConfig(tt.chainID)
			require.Equal(t, tt.want, found)
			require.Equal(t, tt.want, !btcCfg.Empty())
		})
	}
}

func Test_StringMasked(t *testing.T) {
	// create config with defaults
	cfg := config.New(true)

	// mask the config JSON string
	masked := cfg.StringMasked()
	require.NotEmpty(t, masked)

	// should contain necessary fields
	require.Contains(t, masked, "EVMChainConfigs")
	require.Contains(t, masked, "BTCChainConfigs")

	// should not contain endpoint
	require.NotContains(t, masked, "http")
}
