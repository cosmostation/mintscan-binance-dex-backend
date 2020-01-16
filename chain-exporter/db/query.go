package db

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
)

// QueryLatestBlockHeight queries latest block height in database
func (db *Database) QueryLatestBlockHeight() (int32, error) {
	var block schema.BlockInfo
	err := db.Model(&block).
		Order("height DESC").
		Limit(1).
		Select()
	if err != nil {
		return -1, err
	}

	fmt.Println(block)
	fmt.Println(err)

	return block.Height, nil
}
