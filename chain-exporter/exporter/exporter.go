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
	db := db.Connect(&cfg)

	return Exporter{client, db}
}

// StartSyncing starts syncing blockchain data on the active chain
func (ex *Exporter) StartSyncing() {
	height, _ := ex.client.LatestBlockHeight()

	fmt.Println("height: ", height)
}

// cp, err := client.New(cfg.RPCNode, cfg.ClientNode)
// if err != nil {
// 	return errors.Wrap(err, "failed to start RPC client")
// }

// defer cp.Stop() // nolint: errcheck

// db, err := db.OpenDB(cfg)
// if err != nil {
// 	return errors.Wrap(err, "failed to open database connection")
// }

// defer db.Close()
