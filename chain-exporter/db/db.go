package db

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	// columnLength is the column length of varchar type in every table.
	// This needs to be considered again to set it to what specific length is needed, but right now set it to 99999.
	columnLength = 99999
)

const (
	// Define PostgreSQL database indexes to improve the speed of data retrieval operations on a database tables.
	indexBlockHeight            = "CREATE INDEX block_height_idx ON block USING btree(height);"
	indexValidatorConsensusAddr = "CREATE INDEX validator_consensus_address_idx ON validator USING btree(consensus_address);"
	indexPrecommitHeight        = "CREATE INDEX pre_commit_height_idx ON pre_commit USING btree(height);"
	indexPrecommitValidatorAddr = "CREATE INDEX pre_commit_validator_address_idx ON pre_commit USING btree(validator_address);"
	indexTransactionHeight      = "CREATE INDEX transaction_height_idx ON transaction USING btree(height);"
	indexTransactionHash        = "CREATE INDEX transaction_tx_hash_idx ON transaction USING btree(tx_hash);"
	indexTransactionMsgSymbol   = "CREATE INDEX transaction_messages_symbol_idx ON validator USING btree((messages->0->'value'->>'symbol'));"
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
	for _, model := range []interface{}{(*schema.Block)(nil), (*schema.PreCommit)(nil), (*schema.Transaction)(nil),
		(*schema.Validator)(nil), (*schema.StatAssetInfoList1H)(nil), (*schema.StatAssetInfoList24H)(nil)} {

		// Disable pluralization
		orm.SetTableNameInflector(func(s string) string {
			return s
		})

		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     columnLength, // replaces data type from text to varchar type length.
		})

		if err != nil {
			return err
		}
	}

	// Create table indexes and roll back if any index creation fails.
	err := db.createIndexes()
	if err != nil {
		return err
	}

	return nil
}

// createIndexes uses RunInTransaction to run a function in a transaction.
// if function returns an error, transaction is rollbacked, otherwise transaction is committed.
// Create B-Tree indexes to reduce the cost of lookup queries
func (db *Database) createIndexes() error {
	db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := db.Model(schema.Block{}).Exec(indexBlockHeight)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.Validator{}).Exec(indexValidatorConsensusAddr)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.PreCommit{}).Exec(indexPrecommitHeight)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.PreCommit{}).
			Exec(indexPrecommitValidatorAddr)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.Transaction{}).Exec(indexTransactionHeight)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.Transaction{}).Exec(indexTransactionHash)
		if err != nil {
			return err
		}
		_, err = db.Model(schema.Transaction{}).Exec(indexTransactionMsgSymbol)
		if err != nil {
			return err
		}

		return nil
	})

	return nil
}

// --------------------
// Query
// --------------------

// QueryLatestBlockHeight queries latest block height in database.
func (db *Database) QueryLatestBlockHeight() (int64, error) {
	var block schema.Block
	err := db.Model(&block).
		Order("height DESC").
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return 0, err
	}

	if err != nil {
		return -1, err
	}

	return block.Height, nil
}

// QueryValidatorMoniker returns validator's moniker.
func (db *Database) QueryValidatorMoniker(valAddr string) string {
	var validator schema.Validator
	_ = db.Model(&validator).
		Where("consensus_address = ?", valAddr).
		Select()

	return validator.Moniker
}

// --------------------
// Exist
// --------------------

// ExistValidator returns boolean after checking if a validator exists in database.
func (db *Database) ExistValidator(valAddr string) (bool, error) {
	var validator schema.Validator
	ok, err := db.Model(&validator).
		Where("consensus_address = ?", valAddr).
		Exists()

	if err != nil {
		return ok, err
	}

	return ok, nil
}

// --------------------
// Insert
// --------------------

// InsertExportedData inserts exported block, transaction data
// RunInTransaction runs a function in a transaction.
// if function returns an error transaction is rollbacked, otherwise transaction is committed.
func (db *Database) InsertExportedData(block *schema.Block, txs []*schema.Transaction,
	vals []*schema.Validator, precommits []*schema.PreCommit) error {

	err := db.RunInTransaction(func(tx *pg.Tx) error {
		err := tx.Insert(block)
		if err != nil {
			return fmt.Errorf("failed to insert block: %s", err)
		}

		if len(txs) > 0 {
			err := tx.Insert(&txs)
			if err != nil {
				return fmt.Errorf("failed to insert transactions: %s", err)
			}
		}

		if len(vals) > 0 {
			err := tx.Insert(&vals)
			if err != nil {
				return fmt.Errorf("failed to insert validators: %s", err)
			}
		}

		if len(precommits) > 0 {
			err := tx.Insert(&precommits)
			if err != nil {
				return fmt.Errorf("failed to insert precommits: %s", err)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil

}
