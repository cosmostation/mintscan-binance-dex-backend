package exporter

import (
	"fmt"
	"log"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/db"

	"github.com/pkg/errors"

	amino "github.com/tendermint/go-amino"
)

// Exporter wraps the required params to export blockchain
type Exporter struct {
	cdc    *amino.Codec
	client client.Client
	db     *db.Database
}

// NewExporter returns Exporter
func NewExporter() Exporter {
	cfg := config.ParseConfig()
	client := client.NewClient(cfg.Node.RPCNode, cfg.Node.LCDEndpoint)
	db := db.Connect(cfg.DB)

	// Create database tables
	db.CreateTables() // TODO: handle index already exists error

	return Exporter{codec.Codec, client, db}
}

// Start creates database tables and indexes using Postgres ORM library go-pg and
// starts syncing blockchain.
func (ex *Exporter) Start() error {
	go func() {
		for {
			fmt.Println("start - sync blockchain")
			err := ex.sync()
			if err != nil {
				fmt.Printf("error - sync blockchain: %v\n", err)
			}
			fmt.Println("finish - sync blockchain")
			time.Sleep(time.Second)
		}
	}()

	for {
		select {}
	}
}

// sync compares block height between the height saved in your database and
// latest block height on the active chain and calls process to start ingesting blocks.
func (ex *Exporter) sync() error {
	// Query latest block height that is saved in your database
	// Synchronizing blocks from the scratch will return 0 and will ingest accordingly.
	dbHeight, err := ex.db.QueryLatestBlockHeight()
	if dbHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height from database."))
	}

	// Query latest block height on the active network
	latestBlockHeight, err := ex.client.LatestBlockHeight()
	if latestBlockHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height on the active network."))
	}

	// skip the first block since it has no pre-commits
	if dbHeight == 0 {
		dbHeight = 1
	}

	// Ingest all blocks up to the best height
	for i := dbHeight + 1; i <= latestBlockHeight; i++ {
		err = ex.process(i)
		if err != nil {
			return err
		}
		fmt.Printf("synced block %d/%d \n", i, latestBlockHeight)
	}

	return nil
}

func (ex *Exporter) process(height int64) error {
	block, err := ex.client.Block(height)
	if err != nil {
		return fmt.Errorf("failed to query block using rpc client: %t", err)
	}

	prevBlock, err := ex.client.Block(block.Block.LastCommit.Height())
	if err != nil {
		return fmt.Errorf("failed to query block using rpc client: %t", err)
	}

	vals, err := ex.client.Validators(block.Block.LastCommit.Height())
	if err != nil {
		return fmt.Errorf("failed to query validators using rpc client: %t", err)
	}

	resultBlock, err := ex.getBlock(block) // TODO: Fees and RewardsTo Addr
	if err != nil {
		return fmt.Errorf("failed to get block: %t", err)
	}

	resultTxs, err := ex.getTxs(block)
	if err != nil {
		return fmt.Errorf("failed to get transactions: %t", err)
	}

	resultValidators, err := ex.getValidators(prevBlock, vals)
	if err != nil {
		return fmt.Errorf("failed to get validators: %t", err)
	}

	resultPreCommits, err := ex.getPreCommits(block.Block.LastCommit, vals)
	if err != nil {
		return fmt.Errorf("failed to get precommits: %t", err)
	}

	err = ex.db.InsertExportedData(resultBlock, resultTxs, resultValidators, resultPreCommits)
	if err != nil {
		return fmt.Errorf("failed to insert exporterd data: %t", err)
	}

	return nil
}
