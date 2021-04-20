package exporter

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/client"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/db"
)

var (
	// Version is this application's version.
	Version = "dev"

	// Commit is this application's commit hash.
	Commit = ""
)

// Exporter wraps the required params to export blockchain
type Exporter struct {
	l             log.Logger
	client        *client.Client
	db            *db.Database
	ignoreLogs    bool
	genesisHeight int
	allowSyncGap  bool
}

// NewExporter returns Exporter
func NewExporter() *Exporter {
	// Parse config from configuration file (config.yaml).
	config := config.ParseConfig()

	sdkConfig := sdk.GetConfig()
	ctypes.SetBech32Prefixes(sdkConfig)
	ctypes.SetBip44CoinType(sdkConfig)

	// Create new client with node configruation.
	client := client.NewClient(config.Node)

	// Create connection with PostgreSQL database and
	// Ping database to verify connection is success.
	db := db.Connect(config.DB)
	err := db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping database."))
	}

	// Create database tables if not exist already
	db.CreateTables()

	return &Exporter{
		l:             log.DefaultLogger,
		client:        client,
		db:            db,
		ignoreLogs:    config.Processing.IgnoreLogs,
		genesisHeight: config.Processing.GenesisHeight,
		allowSyncGap:  config.Processing.AllowSyncGap,
	}
}

// Start starts to synchronize chain data.
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
	latestBlockHeight, err := ex.client.GetLatestBlockHeight()
	if latestBlockHeight == -1 {
		log.Fatal(errors.Wrap(err, "failed to query the latest block height on the active network"))
	}

	if ex.genesisHeight > 0 {
		if dbHeight == 0 || ex.allowSyncGap {
			dbHeight = int64(ex.genesisHeight)
		}
	}

	// Ingest all blocks up to the latest height
	for i := dbHeight + 1; i <= latestBlockHeight; i++ {
		err = ex.process(i, ex.ignoreLogs)
		if err != nil {
			return err
		}
		ex.l.Printf("synced block %d/%d \n", i, latestBlockHeight)
	}

	return nil
}

// process ingests chain data, such as block, transaction, validator set information
// and save them in database
func (ex *Exporter) process(height int64, ignoreLogs bool) error {
	block, err := ex.client.GetBlock(height)
	if err != nil {
		return fmt.Errorf("failed to query block using rpc client: %s", err)
	}

	var valSet *tmctypes.ResultValidators
	if lastHeight := block.Block.LastCommit.GetHeight(); lastHeight > 0 {
		valSet, err = ex.client.GetValidatorSet(lastHeight)
		if err != nil {
			return fmt.Errorf("failed to query validator set using rpc client: %s", err)
		}
	} else {
		// failed to query validator set using rpc client: RPC error -32603
		// Internal error: height must be greater than 0, but got 0
	}

	vals, err := ex.client.GetValidators()
	if err != nil {
		return fmt.Errorf("failed to query validators using rpc client: %s", err)
	}

	// TODO: Reward Fees Calculation
	resultBlock, err := ex.getBlock(block)
	if err != nil {
		return fmt.Errorf("failed to get block: %s", err)
	}

	resultTxs, err := ex.getTxs(block, ignoreLogs)
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

	// v, _ := json.MarshalIndent(resultBlock, "", "\t")
	// fmt.Println("Result Block:", string(v))
	// v, _ = json.MarshalIndent(resultTxs, "", "\t")
	// fmt.Println("Result Txns:", string(v))
	// v, _ = json.MarshalIndent(resultValidators, "", "\t")
	// fmt.Println("Result Validators:", string(v))
	// v, _ = json.MarshalIndent(resultPreCommits, "", "\t")
	// fmt.Println("Result Pre-Commits:", string(v))

	err = ex.db.InsertExportedData(resultBlock, resultTxs, resultValidators, resultPreCommits)
	if err != nil {
		return fmt.Errorf("failed to insert exported data: %s", err)
	}

	return nil
}
