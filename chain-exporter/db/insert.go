package db

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
	"github.com/go-pg/pg"
)

// InsertExportedData inserts exported block, transaction data
// RunInTransaction runs a function in a transaction.
// if function returns an error transaction is rollbacked, otherwise transaction is committed.
func (db *Database) InsertExportedData(block []*schema.Block, txs []*schema.Transaction,
	vals []*schema.Validator, precommits []*schema.PreCommit) error {

	err := db.RunInTransaction(func(tx *pg.Tx) error {
		if len(block) > 0 {
			err := tx.Insert(&block)
			if err != nil {
				return err
			}
		}

		if len(txs) > 0 {
			err := tx.Insert(&txs)
			if err != nil {
				return err
			}
		}

		if len(vals) > 0 {
			err := tx.Insert(&vals)
			if err != nil {
				return err
			}
		}

		if len(precommits) > 0 {
			err := tx.Insert(&precommits)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
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
