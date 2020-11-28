package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	config := ParseConfig()

	require.NotEmpty(t, config.Node.RPCNode, "Node RPCNode should not be empty")
	require.NotEmpty(t, config.Node.GRPCNode, "Node GRPCNode should not be empty")
	require.NotEmpty(t, config.DB.Host, "Database Host should not be empty")
	require.NotEmpty(t, config.DB.Port, "Database Port should not be empty")
	require.NotEmpty(t, config.DB.User, "Database User should not be empty")
	require.NotEmpty(t, config.DB.Password, "Database Password should not be empty")
	require.NotEmpty(t, config.DB.Database, "Database Name should not be empty")
}
