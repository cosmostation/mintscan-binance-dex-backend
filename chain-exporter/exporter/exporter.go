package exporter

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/db"

	"github.com/pkg/errors"

	amino "github.com/tendermint/go-amino"
)

var (
	// Version is this application's version.
	Version = "Development"

	// Commit is this application's commit hash.
	Commit = ""
)

// Exporter wraps the required params to export blockchain
type Exporter struct {
	l      *log.Logger
	cdc    *amino.Codec
	client *client.Client
	db     *db.Database
}

// NewExporter returns Exporter
func NewExporter() *Exporter {
	l := log.New(os.Stdout, "Chain Exporter ", log.Lshortfile|log.LstdFlags) // [TODO] Project Version

	cfg := config.ParseConfig()

	client := client.NewClient(cfg.Node)

	db := db.Connect(cfg.DB)

	err := db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping database."))
	}

	db.CreateTables()

	return &Exporter{
		l,
		codec.Codec,
		client,
		db,
	}
}

// Start starts to synchronize blockchain data
func (ex *Exporter) Start() error {
	ex.l.Println("Starting Chain Exporter...")
	ex.l.Printf("Version: %s | Commit Hash: %s", Version, Commit)

	go func() {
		for {
			ex.l.Println("start - sync blockchain")
			err := ex.sync()
			if err != nil {
				ex.l.Printf("error - sync blockchain: %v\n", err)
			}
			ex.l.Println("finish - sync blockchain")
			time.Sleep(time.Second)
		}
	}()

	for {
		select {}
	}
}

// sync compares block height between the height saved in your database and
// the latest block height on the active chain and calls process to start ingesting data.
func (ex *Exporter) sync() error {
	// Query latest block height saved in database
	dbHeight, err := ex.db.QueryLatestBlockHeight()
	if dbHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height saved in database"))
	}

	// Query latest block height on the active network
	latestBlockHeight, err := ex.client.LatestBlockHeight()
	if latestBlockHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height on the active network"))
	}

	// Synchronizing blocks from the scratch will return 0 and will ingest accordingly.
	// Skip the first block since it has no pre-commits
	if dbHeight == 0 {
		dbHeight = 1
	}

	// Ingest all blocks up to the latest height
	for i := dbHeight + 1; i <= latestBlockHeight; i++ {
		err = ex.process(i)
		if err != nil {
			return err
		}
		ex.l.Printf("synced block %d/%d \n", i, latestBlockHeight)
	}

	return nil
}

// process ingests chain data, such as block, transaction, validator set information
// and save them in database
func (ex *Exporter) process(height int64) error {
	block, err := ex.client.Block(height)
	if err != nil {
		return fmt.Errorf("failed to query block using rpc client: %s", err)
	}

	valSet, err := ex.client.ValidatorSet(block.Block.LastCommit.Height())
	if err != nil {
		return fmt.Errorf("failed to query validator set using rpc client: %s", err)
	}

	vals, err := ex.client.Validators()
	if err != nil {
		return fmt.Errorf("failed to query validators using rpc client: %s", err)
	}

	resultBlock, err := ex.getBlock(block) // TODO: Reward Fees Calculation
	if err != nil {
		return fmt.Errorf("failed to get block: %s", err)
	}

	resultTxs, err := ex.getTxs(block)
	if err != nil {
		return fmt.Errorf("failed to get transactions: %s", err)
	}

	resultValidators, err := ex.getValidators(vals)
	if err != nil {
		return fmt.Errorf("failed to get validators: %s", err)
	}

	resultPreCommits, err := ex.getPreCommits(block.Block.LastCommit, valSet)
	if err != nil {
		return fmt.Errorf("failed to get precommits: %s", err)
	}

	err = ex.db.InsertExportedData(resultBlock, resultTxs, resultValidators, resultPreCommits)
	if err != nil {
		return fmt.Errorf("failed to insert exporterd data: %s", err)
	}

	return nil
}
