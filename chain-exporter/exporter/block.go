package exporter

import (
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/schema"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// getBlock exports block information.
func (ex *Exporter) getBlock(block *tmctypes.ResultBlock) (*schema.Block, error) {
	b := schema.NewBlock(schema.Block{
		Height:        block.Block.Height,
		Proposer:      block.Block.ProposerAddress.String(),
		Moniker:       ex.db.QueryValidatorMoniker(block.Block.ProposerAddress.String()),
		BlockHash:     block.BlockID.Hash.String(),
		ParentHash:    block.Block.Header.LastBlockID.Hash.String(),
		NumPrecommits: int64(len(block.Block.LastCommit.Signatures)),
		NumTxs:        int64(len(block.Block.Txs)),
		Timestamp:     block.Block.Time,
	})

	return b, nil
}
