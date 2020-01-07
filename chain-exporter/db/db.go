package db

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Database implements a wrapper of pg.DB
type Database struct {
	*pg.DB
}

// Connect opens a database connections with the given database connection info from config.
// It returns a database connection handle or an error if the connection fails.
func Connect(cfg *config.Config) *Database {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.DB.Host + ":" + string(cfg.DB.Port),
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Database: cfg.DB.Name,
	})

	return &Database{db}
}

// CreateSchema creates database tables using object relational mapping (ORM)
func (db *Database) CreateSchema() error {
	for _, model := range []interface{}{(*schema.BlockInfo)(nil), (*schema.TransactionInfo)(nil)} {

		// Disable pluralization
		orm.SetTableNameInflector(func(s string) string {
			return s
		})

		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     100, // replaces PostgreSQL data type `text` with `varchar(n)`
		})
		if err != nil {
			return err
		}
	}

	// RunInTransaction creates indexes to reduce the cost of lookup queries in case of server traffic jams.
	// If function returns an error transaction is rollbacked, otherwise transaction is committed.
	err := db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := db.Model(schema.BlockInfo{}).Exec(`CREATE INDEX block_info_height_idx ON block_info USING btree(height);`)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.TransactionInfo{}).Exec(`CREATE INDEX transaction_info_height_idx ON transaction_info USING btree(height);`)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
