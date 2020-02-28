package exporter

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// getBlock parses block information and wrap into Block schema struct
func (ex *Exporter) getBlock(block *tmctypes.ResultBlock) ([]*schema.Block, error) {
	blocks := make([]*schema.Block, 0)

	tempBlock := &schema.Block{
		Height:        block.Block.Height,
		Proposer:      block.Block.ProposerAddress.String(),
		Moniker:       ex.db.QueryMoniker(block.Block.ProposerAddress.String()),
		BlockHash:     block.BlockMeta.BlockID.Hash.String(),
		ParentHash:    block.BlockMeta.Header.LastBlockID.Hash.String(),
		NumPrecommits: int64(len(block.Block.LastCommit.Precommits)),
		NumTxs:        block.Block.NumTxs,
		TotalTxs:      block.Block.TotalTxs,
		Timestamp:     block.Block.Time,
	}

	blocks = append(blocks, tempBlock)

	return blocks, nil
}
