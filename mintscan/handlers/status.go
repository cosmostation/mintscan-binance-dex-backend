package handlers

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// Status is a status handler
type Status struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewStatus creates a new Status handler with the given params
func NewStatus(l *log.Logger, client *client.Client, db *db.Database) *Status {
	return &Status{l, client, db}
}

// GetStatus returns current status on the active chain
func (s *Status) GetStatus(rw http.ResponseWriter, r *http.Request) {
	status, err := s.client.Status()
	if err != nil {
		s.l.Printf("failed to query status: %s\n", err)
	}

	validatorSet, err := s.client.ValidatorSet(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		s.l.Printf("failed to query validators et: %s\n", err)
	}

	block, err := s.client.Block(status.SyncInfo.LatestBlockHeight)
	if err != nil {
		s.l.Printf("failed to query block: %s\n", err)
	}

	prevBlock, err := s.client.Block(status.SyncInfo.LatestBlockHeight - 1)
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

	utils.Respond(rw, result)
	return
}
