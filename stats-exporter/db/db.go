package db

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/schema"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Database implements a wrapper of golang ORM with focus on PostgreSQL
type Database struct {
	*pg.DB
}

// Connect opens a database connections with the given database connection info from config.
// It returns a database connection handle or an error if the connection fails.
func Connect(cfg config.DBConfig) *Database {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Table,
	})

	return &Database{db}
}

// Ping returns a database connection handle or an error if the connection fails.
func (db *Database) Ping() error {
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return err
	}

	return nil
}

// CreateTables creates database tables using object relational mapping (ORM)
func (db *Database) CreateTables() error {
	for _, model := range []interface{}{(*schema.StatAssetInfoList1H)(nil), (*schema.StatAssetInfoList24H)(nil)} {

		orm.SetTableNameInflector(func(s string) string { // disable pluralization
			return s
		})

		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     20000, // replaces data type from `text` to `varchar(n)`
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// SaveAssetInfoList1H inserts asset information list every hour
func (db *Database) SaveAssetInfoList1H(assets []*schema.StatAssetInfoList1H) error {
	err := db.Insert(&assets)
	if err != nil {
		return err
	}

	return nil
}

// SaveAssetInfoList24H inserts asset information list every 24 hours
func (db *Database) SaveAssetInfoList24H(assets []*schema.StatAssetInfoList24H) error {
	err := db.Insert(&assets)
	if err != nil {
		return err
	}

	return nil
}
