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

// QueryMoniker queries validator moniker
func (db *Database) QueryMoniker(valAddr string) string {
	var validator schema.Validator
	_ = db.Model(&validator).
		Where("consensus_address = ?", valAddr).
		Select()

	return validator.Moniker
}

// ExistValidator checks to see if a validator exists
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

// ExistToken checks to see if a token exists
func (db *Database) ExistToken(originalSymbol string) (bool, error) {
	var token schema.Token
	ok, err := db.Model(&token).
		Where("original_symbol = ?", originalSymbol).
		Exists()
	if err != nil {
		return ok, err
	}

	return ok, nil
}
