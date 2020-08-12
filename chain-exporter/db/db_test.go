package db

import (
	"os"
	"testing"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"github.com/stretchr/testify/require"
)

var db *Database

func TestMain(m *testing.M) {
	config := config.ParseConfig()

	db = Connect(config.DB)

	os.Exit(m.Run())
}

func TestCreate_Tables(t *testing.T) {
	err := db.Ping()
	require.NoError(t, err)

	tables := []interface{}{
		(*schema.Block)(nil),
		(*schema.PreCommit)(nil),
		(*schema.Transaction)(nil),
		(*schema.Validator)(nil),
		(*schema.StatAssetInfoList1H)(nil),
		(*schema.StatAssetInfoList24H)(nil),
	}

	for _, table := range tables {
		orm.SetTableNameInflector(func(s string) string {
			return s
		})

		err := db.CreateTable(table, &orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     columnLength,
		})

		require.NoError(t, err)
	}
}

func TestConnection(t *testing.T) {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	require.NoError(t, err)

	require.Equal(t, n, 1, "failed to ping database")
}
