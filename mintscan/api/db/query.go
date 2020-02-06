package db

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/schema"

	"github.com/go-pg/pg"
)

// QueryBlocks queries blocks with pagination params, such as limit, before, after, and offset
func (db *Database) QueryBlocks(limit int, before int, after int, offset int) ([]schema.Block, error) {
	blocks := make([]schema.Block, 0)

	var err error

	switch {
	case before > 0:
		err = db.Model(&blocks).
			Where("height < ?", before).
			Limit(limit).
			Order("id DESC").
			Select()
	case after > 0:
		err = db.Model(&blocks).
			Where("height > ?", after).
			Limit(limit).
			Order("id ASC").
			Select()
	case offset >= 0:
		err = db.Model(&blocks).
			Limit(limit).
			Offset(offset).
			Order("id DESC").
			Select()
	}

	if err == pg.ErrNoRows {
		return blocks, fmt.Errorf("no rows in block table: %t", err)
	}

	if err != nil {
		return blocks, fmt.Errorf("unexpected database error: %t", err)
	}

	return blocks, nil
}

// QueryTxs queries transactions with pagination params, such as limit, before, after, and offset
func (db *Database) QueryTxs(limit int, before int, after int, offset int) ([]schema.Transaction, error) {
	txs := make([]schema.Transaction, 0)

	var err error

	switch {
	case before > 0:
		err = db.Model(&txs).
			Where("height < ?", before).
			Limit(limit).
			Order("id DESC").
			Select()
	case after > 0:
		err = db.Model(&txs).
			Where("height > ?", after).
			Limit(limit).
			Order("id ASC").
			Select()
	case offset >= 0:
		err = db.Model(&txs).
			Limit(limit).
			Offset(offset).
			Order("id DESC").
			Select()
	}

	if err == pg.ErrNoRows {
		return txs, fmt.Errorf("no rows in block table: %t", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %t", err)
	}

	return txs, nil
}

// QueryTx queries particular transaction with height
func (db *Database) QueryTx(height int64) ([]schema.Transaction, error) {
	txs := make([]schema.Transaction, 0)

	err := db.Model(&txs).
		Where("height = ?", height).
		Select()

	if err == pg.ErrNoRows {
		return txs, fmt.Errorf("no rows in block table: %t", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %t", err)
	}

	return txs, nil
}
