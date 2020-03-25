package services

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// GetStatus returns current status on the active chain
func GetStatus(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	status, err := c.Status()
	if err != nil {
		log.Printf("failed to query status: %s\n", err)
	}

	validatorSet, err := c.ValidatorSet(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		log.Printf("failed to query validators et: %s\n", err)
	}

	block, err := c.Block(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		log.Printf("failed to query block: %s\n", err)
	}

	prevBlock, err := c.Block(status.SyncInfo.LatestBlockHeight - 1)
	if err != nil {
		log.Printf("failed to query previous block: %s\n", err)
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
