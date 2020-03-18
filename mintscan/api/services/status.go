package services

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

// GetStatus returns current status on the active chain
func GetStatus(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	status, _ := client.Status()
	validatorSet, _ := client.ValidatorSet(status.SyncInfo.LatestBlockHeight)

	block, err := client.Block(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		log.Printf("failed to query block using rpc client: %s\n", err)
	}

	prevBlock, err := client.Block(status.SyncInfo.LatestBlockHeight - 1)
	if err != nil {
		log.Printf("failed to query previous block using rpc client: %s\n", err)
	}

	blockTime := block.Block.Time.UTC().
		Sub(prevBlock.Block.Time.UTC()).Seconds()

	result := &models.Status{
		ChainID:           status.NodeInfo.Network,
		BlockTime:         blockTime,
		LatestBlockHeight: status.SyncInfo.LatestBlockHeight,
		TotalValidatorNum: len(validatorSet.Validators),
		Timestamp:         status.SyncInfo.LatestBlockTime,
	}

	utils.Respond(w, result)
	return nil
}
