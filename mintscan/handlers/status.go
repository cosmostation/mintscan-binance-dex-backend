package handlers

import (
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
)

// GetStatus returns current status on the active chain
func GetStatus(c *gin.Context) {
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

	models.Respond(c.Writer, result)
	return
}
