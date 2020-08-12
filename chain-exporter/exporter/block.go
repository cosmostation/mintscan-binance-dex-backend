package exporter

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// getBlock exports block information.
func (ex *Exporter) getBlock(block *tmctypes.ResultBlock) (*schema.Block, error) {
	b := schema.NewBlock(schema.Block{
		Height:        block.Block.Height,
		Proposer:      block.Block.ProposerAddress.String(),
		Moniker:       ex.db.QueryValidatorMoniker(block.Block.ProposerAddress.String()),
		BlockHash:     block.BlockMeta.BlockID.Hash.String(),
		ParentHash:    block.BlockMeta.Header.LastBlockID.Hash.String(),
		NumPrecommits: int64(len(block.Block.LastCommit.Precommits)),
		NumTxs:        block.Block.NumTxs,
		TotalTxs:      block.Block.TotalTxs,
		Timestamp:     block.Block.Time,
	})

	return b, nil
}
