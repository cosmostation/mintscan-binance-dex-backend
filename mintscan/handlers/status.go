package handlers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
)

// GetStatus returns current status on the active chain
func GetStatus(rw http.ResponseWriter, r *http.Request) {
	status, err := s.client.GetStatus()
	if err != nil {
		s.l.Printf("failed to query status: %s\n", err)
	}

	validatorSet, err := s.client.GetValidatorSet(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		s.l.Printf("failed to query validators et: %s\n", err)
	}

	block, err := s.client.GetBlock(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		s.l.Printf("failed to query block: %s\n", err)
	}

	prevBlock, err := s.client.GetBlock(status.SyncInfo.LatestBlockHeight - 1)
	if err != nil {
		s.l.Printf("failed to query previous block: %s\n", err)
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

	models.Respond(rw, result)
	return
}
