package schema

import "time"

// Block defines the schema for block information
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
