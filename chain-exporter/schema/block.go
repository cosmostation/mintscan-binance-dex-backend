package schema

import "time"

// BlockInfo represents the information a block contains
type BlockInfo struct {
	ID        int32     `json:"id" sql:",pk"`
	Height    int32     `json:"height" sql:",notnull"`
	Proposer  string    `json:"proposer" sql:",notnull"`
	BlockHash string    `json:"block_hash" sql:",notnull"`
	Precommit string    `json:"pre_commit"`
	NumTxs    int32     `json:"num_txs" sql:"default:0"`
	TotalTxs  int32     `json:"total_txs" sql:"default:0"`
	Timestamp time.Time `json:"timestamp" sql:"default:now()"`
}
