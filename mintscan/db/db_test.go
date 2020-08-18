package db

import (
	"os"
	"testing"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"

	"github.com/stretchr/testify/require"

	"github.com/go-pg/pg"
)

var db *Database

func TestMain(m *testing.M) {
	config := config.ParseConfig()
	db = Connect(config.DB)

	os.Exit(m.Run())
}

func TestConnection(t *testing.T) {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	require.NoError(t, err)

	require.Equal(t, n, 1, "failed to ping database")
}
