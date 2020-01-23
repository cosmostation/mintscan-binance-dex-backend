package exporter

import (
	"fmt"
	"log"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/db"
	"github.com/pkg/errors"
)

// Exporter wraps the required params to export blockchain
type Exporter struct {
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

	return Exporter{client, db}
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
		fmt.Printf("failed to get block information: %t\n", err)
		return err
	}

	txs, err := ex.client.Txs(block)
	if err != nil {
		fmt.Printf("failed to get transactions in a block: %t\n", err)
		return err
	}

	fmt.Println(txs)

	return nil
}
