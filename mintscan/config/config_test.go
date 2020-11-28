package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	config := ParseConfig()

	require.NotEmpty(t, config.Node.RPCNode, "Node RPCNode should not be empty")
	require.NotEmpty(t, config.Node.ExchangeAPIEndpoint, "Node ExchangeAPIEndpoint should not be empty")
	require.NotEmpty(t, config.DB.Host, "Database Host is empty")
	require.NotEmpty(t, config.DB.Port, "Database Port is empty")
	require.NotEmpty(t, config.DB.User, "Database User is empty")
	require.NotEmpty(t, config.DB.Password, "Database Password is empty")
	require.NotEmpty(t, config.DB.Database, "Database Name is empty")
	require.NotEmpty(t, config.Web.Port, "Web Port is empty")
	require.NotEmpty(t, config.Market.CoinGeckoEndpoint, "CoinGecko Endpoint is empty")
}
