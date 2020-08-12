package schema

import "time"

// Block defines the structure for block information.
type Block struct {
	ID            int32     `json:"id" sql:",pk"`
	Height        int64     `json:"height" sql:",notnull"`
	Proposer      string    `json:"proposer" sql:",notnull"`
	Moniker       string    `json:"moniker" sql:",notnull"`
	BlockHash     string    `json:"block_hash" sql:",notnull,unique"`
	ParentHash    string    `json:"parent_hash" sql:",notnull"`
	NumPrecommits int64     `json:"num_pre_commits" sql:",notnull"`
	NumTxs        int64     `json:"num_txs" sql:"default:0"`
	TotalTxs      int64     `json:"total_txs" sql:"default:0"`
	Timestamp     time.Time `json:"timestamp" sql:"default:now()"`
}

// NewBlock returns a new Block.
func NewBlock(b Block) *Block {
	return &Block{
		Height:        b.Height,
		Proposer:      b.Proposer,
		Moniker:       b.Moniker,
		BlockHash:     b.BlockHash,
		ParentHash:    b.ParentHash,
		NumPrecommits: b.NumPrecommits,
		NumTxs:        b.NumTxs,
		TotalTxs:      b.TotalTxs,
		Timestamp:     b.Timestamp,
	}
}
