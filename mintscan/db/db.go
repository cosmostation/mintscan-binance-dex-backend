package db

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"

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
	orm.SetTableNameInflector(func(s string) string { // disable pluralization
		return s
	})

	db := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Table,
	})

	return &Database{db}
}

// Ping returns a database connection handle
// An error is returned if the connection fails.
func (db *Database) Ping() error {
	_, err := db.Exec("SELECT 1")
	if err != nil {
		return err
	}

	return nil
}
