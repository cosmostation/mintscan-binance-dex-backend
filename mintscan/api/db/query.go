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

// QueryLatestBlockHeight queries latest block height saved in database
func (db *Database) QueryLatestBlockHeight() (int64, error) {
	var block schema.Block

	err := db.Model(&block).
		Column("height").
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return 0, fmt.Errorf("no rows in block table: %t", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %t", err)
	}

	return block.Height, nil
}

// QueryTotalTxsNum queries total number of transactions
func (db *Database) QueryTotalTxsNum(height int64) (int64, error) {
	var block schema.Block

	err := db.Model(&block).
		Where("height = ?", height).
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return 0, fmt.Errorf("no rows in block table: %t", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %t", err)
	}

	return block.TotalTxs, nil
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

// QueryTxsByType queries transactions with tx type and start and end time
func (db *Database) QueryTxsByType(txType string, startTime int64, endTime int64, limit int, before int) ([]schema.Transaction, error) {
	txs := make([]schema.Transaction, 0)

	err := db.Model(&txs).
		Where("(messages->0->>'type' = ?) AND TIMESTAMP BETWEEN TO_TIMESTAMP(?) AND TO_TIMESTAMP(?) AND height < ?", txType, startTime, endTime, before).
		Limit(limit).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return txs, fmt.Errorf("no rows in block table: %t", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %t", err)
	}

	return txs, nil
}
