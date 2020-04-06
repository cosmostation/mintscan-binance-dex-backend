package db

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/schema"

	"github.com/go-pg/pg"
)

// QueryBlocks queries blocks with pagination params, such as limit, before, after, and offset
func (db *Database) QueryBlocks(before int, after int, limit int) ([]schema.Block, error) {
	blocks := make([]schema.Block, 0)

	var err error

	switch {
	case before > 0:
		err = db.Model(&blocks).
			Where("height < ?", before).
			Limit(limit).
			Order("id DESC").
			Select()
	case after >= 0:
		err = db.Model(&blocks).
			Where("height > ?", after).
			Limit(limit).
			Order("id ASC").
			Select()
	default:
		err = db.Model(&blocks).
			Limit(limit).
			Order("id DESC").
			Select()
	}

	if err == pg.ErrNoRows {
		return blocks, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return blocks, fmt.Errorf("unexpected database error: %s", err)
	}

	return blocks, nil
}

// QueryLatestBlock queries latest block information saved in database
func (db *Database) QueryLatestBlock() (schema.Block, error) {
	var block schema.Block

	err := db.Model(&block).
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return schema.Block{}, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return schema.Block{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return block, nil
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
		return 0, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %s", err)
	}

	return block.Height, nil
}

// QueryTotalTxsNum queries total number of transactions up until that height
func (db *Database) QueryTotalTxsNum(height int64) (int64, error) {
	var block schema.Block

	err := db.Model(&block).
		Where("height = ?", height).
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return 0, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %s", err)
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
		return txs, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %s", err)
	}

	return txs, nil
}

// QueryTxs queries transactions with pagination params, such as limit, before, after, and offset
func (db *Database) QueryTxs(before int, after int, limit int) ([]schema.Transaction, error) {
	txs := make([]schema.Transaction, 0)

	var err error

	switch {
	case before > 0:
		err = db.Model(&txs).
			Where("id < ?", before).
			Limit(limit).
			Order("id DESC").
			Select()
	case after >= 0:
		err = db.Model(&txs).
			Where("id > ?", after).
			Limit(limit).
			Order("id ASC").
			Select()
	default:
		err = db.Model(&txs).
			Limit(limit).
			Order("id DESC").
			Select()
	}

	if err == pg.ErrNoRows {
		return txs, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %s", err)
	}

	return txs, nil
}

// QueryTxByHash queries transaction by transaction hash
func (db *Database) QueryTxByHash(hash string) (schema.Transaction, error) {
	var tx schema.Transaction
	err := db.Model(&tx).
		Where("tx_hash = ?", hash).
		Select()

	if err == pg.ErrNoRows {
		return tx, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return tx, fmt.Errorf("unexpected database error: %s", err)
	}

	return tx, nil
}

// QueryTxsByType queries transactions with tx type and start and end time
func (db *Database) QueryTxsByType(txType string, startTime int64, endTime int64, before int, after int, limit int) ([]schema.Transaction, error) {
	txs := make([]schema.Transaction, 0)

	var err error

	switch {
	case before > 0:
		err = db.Model(&txs).
			Where("(messages->0->>'type' = ?) AND TIMESTAMP BETWEEN TO_TIMESTAMP(?) AND TO_TIMESTAMP(?) AND id < ?", txType, startTime, endTime, before).
			Limit(limit).
			Order("id DESC").
			Select()
	case after >= 0:
		err = db.Model(&txs).
			Where("(messages->0->>'type' = ?) AND TIMESTAMP BETWEEN TO_TIMESTAMP(?) AND TO_TIMESTAMP(?) AND id > ?", txType, startTime, endTime, after).
			Limit(limit).
			Order("id ASC").
			Select()
	default:
		err = db.Model(&txs).
			Where("(messages->0->>'type' = ?) AND TIMESTAMP BETWEEN TO_TIMESTAMP(?) AND TO_TIMESTAMP(?) AND id < ?", txType, startTime, endTime, before).
			Limit(limit).
			Order("id DESC").
			Select()
	}

	if err == pg.ErrNoRows {
		return txs, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %s", err)
	}

	return txs, nil
}

// ExistToken checks to see if a token exists
func (db *Database) ExistToken(originalSymbol string) (bool, error) {
	var token models.Token
	ok, err := db.Model(&token).
		Where("original_symbol = ?", originalSymbol).
		Exists()
	if err != nil {
		return ok, err
	}

	return ok, nil
}

// CountTotalTxsNum counts total number of transactions saved in transaction table
// Note that count(*) raises performance issue as more txs saved in database
func (db *Database) CountTotalTxsNum() (int32, error) {
	var tx schema.Transaction

	err := db.Model(&tx).
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return 0, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %s", err)
	}

	return tx.ID, nil
}

// QueryAssetChartHistory queries asset chart history
// Stats Exporter needs to be executed and run at least 24 hours to get the result
func (db *Database) QueryAssetChartHistory(asset string, limit int) ([]schema.StatAssetInfoList1H, error) {
	chartHistory := make([]schema.StatAssetInfoList1H, 0)

	err := db.Model(&chartHistory).
		Where("asset = ?", asset).
		Limit(limit).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return chartHistory, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return chartHistory, fmt.Errorf("unexpected database error: %s", err)
	}

	return chartHistory, nil
}

// QueryValidators queries validators in a validator set saved in database
func (db *Database) QueryValidators() ([]*schema.Validator, error) {
	vals := make([]*schema.Validator, 0)

	err := db.Model(&vals).
		Order("tokens DESC").
		Select()

	if err == pg.ErrNoRows {
		return vals, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return vals, fmt.Errorf("unexpected database error: %s", err)
	}

	return vals, nil
}

// QueryValidatorByOperAddr queries validators in a validator set saved in database
func (db *Database) QueryValidatorByOperAddr(address string) (schema.Validator, error) {
	var val schema.Validator

	err := db.Model(&val).
		Where("operator_address = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return val, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return val, fmt.Errorf("unexpected database error: %s", err)
	}

	return val, nil
}

// QueryValidatorByAccountAddr queries validators in a validator set saved in database
func (db *Database) QueryValidatorByAccountAddr(address string) (schema.Validator, error) {
	var val schema.Validator

	err := db.Model(&val).
		Where("account_address = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return val, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return val, fmt.Errorf("unexpected database error: %s", err)
	}

	return val, nil
}

// QueryValidatorByConsAddr queries validators in a validator set saved in database
func (db *Database) QueryValidatorByConsAddr(address string) (schema.Validator, error) {
	var val schema.Validator

	err := db.Model(&val).
		Where("consensus_address = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return val, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return val, fmt.Errorf("unexpected database error: %s", err)
	}

	return val, nil
}

// QueryValidatorByMoniker queries validators in a validator set saved in database
func (db *Database) QueryValidatorByMoniker(address string) (schema.Validator, error) {
	var val schema.Validator

	err := db.Model(&val).
		Where("moniker = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return val, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return val, fmt.Errorf("unexpected database error: %s", err)
	}

	return val, nil
}
