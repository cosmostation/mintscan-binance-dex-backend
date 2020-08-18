package db

import (
	"fmt"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/schema"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Database implements a wrapper of golang ORM with focus on PostgreSQL.
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
		Addr:         cfg.Host + ":" + cfg.Port,
		User:         cfg.User,
		Password:     cfg.Password,
		Database:     cfg.Table,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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

// --------------------
// Query
// --------------------

// QueryBlocks queries blocks with given params and return them.
func (db *Database) QueryBlocks(before int, after int, limit int) (blocks []schema.Block, err error) {
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
		return []schema.Block{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return []schema.Block{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return blocks, nil
}

// QueryLatestBlock queries latest block information saved in database.
func (db *Database) QueryLatestBlock() (block schema.Block, err error) {
	err = db.Model(&block).
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return schema.Block{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return schema.Block{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return block, nil
}

// QueryLatestBlockHeight queries latest block height saved in database and return it.
func (db *Database) QueryLatestBlockHeight() (int64, error) {
	var block schema.Block
	err := db.Model(&block).
		Column("height").
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return 0, fmt.Errorf("found no rows in block table: %s", err)
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
		return 0, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %s", err)
	}

	return block.TotalTxs, nil
}

// QueryTx queries particular transaction with height
func (db *Database) QueryTx(height int64) (txs []schema.Transaction, err error) {
	err = db.Model(&txs).
		Where("height = ?", height).
		Select()

	if err == pg.ErrNoRows {
		return txs, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return txs, fmt.Errorf("unexpected database error: %s", err)
	}

	return txs, nil
}

// QueryTxs queries transactions with given params and return them.
func (db *Database) QueryTxs(before int, after int, limit int) (txs []schema.Transaction, err error) {
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
		return []schema.Transaction{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return []schema.Transaction{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return txs, nil
}

// QueryTxByHash queries transaction by transaction hash and return it.
func (db *Database) QueryTxByHash(hash string) (tx schema.Transaction, err error) {
	err = db.Model(&tx).
		Where("tx_hash = ?", hash).
		Select()

	if err == pg.ErrNoRows {
		return schema.Transaction{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return schema.Transaction{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return tx, nil
}

// QueryTxsByType queries transactions with tx type, start time and end time and return them.
func (db *Database) QueryTxsByType(txType string, startTime int64, endTime int64, before int, after int, limit int) (txs []schema.Transaction, err error) {
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
		return []schema.Transaction{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return []schema.Transaction{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return txs, nil
}

// QueryAssetChartHistory queries asset chart history
// Stats Exporter needs to be executed and run at least 24 hours to get the result
func (db *Database) QueryAssetChartHistory(asset string, limit int) (stats []schema.StatAssetInfoList1H, err error) {
	err = db.Model(&stats).
		Where("asset = ?", asset).
		Limit(limit).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return []schema.StatAssetInfoList1H{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return []schema.StatAssetInfoList1H{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return stats, nil
}

// QueryValidators queries validators in a validator set saved in database
func (db *Database) QueryValidators() (validators []*schema.Validator, err error) {
	err = db.Model(&validators).
		Order("tokens DESC").
		Select()

	if err == pg.ErrNoRows {
		return []*schema.Validator{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return []*schema.Validator{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return validators, nil
}

// QueryValidatorByOperAddr queries validators in a validator set saved in database
func (db *Database) QueryValidatorByOperAddr(address string) (schema.Validator, error) {
	var val schema.Validator

	err := db.Model(&val).
		Where("operator_address = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return val, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return val, fmt.Errorf("unexpected database error: %s", err)
	}

	return val, nil
}

// QueryValidatorByAccountAddr queries validators in a validator set saved in database.
func (db *Database) QueryValidatorByAccountAddr(address string) (validator schema.Validator, err error) {
	err = db.Model(&validator).
		Where("account_address = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return schema.Validator{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return schema.Validator{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return validator, nil
}

// QueryValidatorByConsAddr queries validators in a validator set saved in database.
func (db *Database) QueryValidatorByConsAddr(address string) (validator schema.Validator, err error) {
	err = db.Model(&validator).
		Where("consensus_address = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return schema.Validator{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return schema.Validator{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return validator, nil
}

// QueryValidatorByMoniker queries validators in a validator set saved in database
func (db *Database) QueryValidatorByMoniker(address string) (validator schema.Validator, err error) {
	err = db.Model(&validator).
		Where("moniker = ?", address).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return schema.Validator{}, fmt.Errorf("found no rows in block table: %s", err)
	}

	if err != nil {
		return schema.Validator{}, fmt.Errorf("unexpected database error: %s", err)
	}

	return validator, nil
}

// --------------------
// Count
// --------------------

// CountTotalTxsNum counts total number of transactions saved in transaction table
// Note that count(*) raises performance issue as more txs saved in database
func (db *Database) CountTotalTxsNum() (int32, error) {
	var tx schema.Transaction
	err := db.Model(&tx).
		Limit(1).
		Order("id DESC").
		Select()

	if err == pg.ErrNoRows {
		return -1, fmt.Errorf("no rows in block table: %s", err)
	}

	if err != nil {
		return 0, fmt.Errorf("unexpected database error: %s", err)
	}

	return tx.ID, nil
}

// --------------------
// Exist
// --------------------

// ExistToken checks to see if a token exists.
func (db *Database) ExistToken(originalSymbol string) (exist bool, err error) {
	var token models.Token
	exist, err = db.Model(&token).
		Where("original_symbol = ?", originalSymbol).
		Exists()

	if err != nil {
		return false, err
	}

	return exist, nil
}
