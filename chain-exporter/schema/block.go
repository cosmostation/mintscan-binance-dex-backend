package schema

import "time"

// BlockInfo represents the information a block contains
type BlockInfo struct {
	ID           int32     `json:"id" sql:",pk"`
	Height       int64     `json:"height" sql:",notnull"`
	Proposer     string    `json:"proposer" sql:",notnull"`
	BlockHash    string    `json:"block_hash" sql:",notnull,unique"`
	NumPrecommit string    `json:"num_pre_commit" sql:",notnull"`
	NumTxs       int64     `json:"num_txs" sql:"default:0"`
	TotalTxs     int64     `json:"total_txs" sql:"default:0"`
	Timestamp    time.Time `json:"timestamp" sql:"default:now()"`
}
