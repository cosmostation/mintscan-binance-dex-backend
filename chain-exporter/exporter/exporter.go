package exporter

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/rpc"
)

// Exporter wraps the required params to export blockchain
type Exporter struct {
	client rpc.Client
	db     *db.Database
}

// NewExporter returns Exporter
func NewExporter() Exporter {
	cfg := config.ParseConfig()
	client := rpc.NewClient(cfg.Node.RPCNode, cfg.Node.LCDEndpoint)
	db := db.Connect(cfg.DB)

	return Exporter{client, db}
}

// StartSyncing compares block height between the height saved in database and
// latest block height on the active chain and starts syncing blockchain.
func (ex *Exporter) StartSyncing() error {
	// Create database schemas
	err := ex.db.CreateTables()
	if err != nil {
		return fmt.Errorf("failed to create database tables: %t", err)
	}

	dbHeight, err := ex.db.QueryLatestBlockHeight()
	if err != nil {
		return fmt.Errorf("failed to query latest block height in database: %t", err)
	}

	latestBlockHeight, _ := ex.client.LatestBlockHeight()

	fmt.Println("dbHeight: ", dbHeight)
	fmt.Println("latestBlockHeight: ", latestBlockHeight)

	return nil
}

func (ex *Exporter) process() {

}

//
//
//

type Queue chan string

func NewQueue() Queue {
	return make(chan string)
}
