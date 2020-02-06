package db

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	"github.com/go-pg/pg"
)

/*
	For any type of database errors, return -1
*/

// QueryLatestBlockHeight queries latest block height in database
func (db *Database) QueryLatestBlockHeight() (int64, error) {
	var block schema.Block
	err := db.Model(&block).
		Order("height DESC").
		Limit(1).
		Select()

	// return 0 when there is no row in result set
	if err == pg.ErrNoRows {
		return 0, err
	}

	if err != nil {
		return -1, err
	}

	return block.Height, nil
}

// ExistValidator checks to see if a validator exists
func (db *Database) ExistValidator(valAddr string) (bool, error) {
	var validator schema.Validator
	ok, err := db.Model(&validator).
		Where("validator_address = ?", valAddr).
		Exists()

	if err != nil {
		return ok, err
	}

	return ok, nil
}
