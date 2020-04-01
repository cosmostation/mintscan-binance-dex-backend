package models

import "time"

type (
	// ResultBlocks defines the structure for block result response
	ResultBlocks struct {
		Paging Paging      `json:"paging"`
		Data   []BlockData `json:"data"`
	}

	// BlockData wraps block data
	BlockData struct {
		Height        int64     `json:"height"`
		Proposer      string    `json:"proposer"`
		Moniker       string    `json:"moniker"`
		BlockHash     string    `json:"block_hash"`
		ParentHash    string    `json:"parent_hash"`
		NumPrecommits int64     `json:"num_pre_commits" sql:",notnull"`
		NumTxs        int64     `json:"num_txs" sql:"default:0"`
		TotalTxs      int64     `json:"total_txs" sql:"default:0"`
		Txs           []Txs     `json:"txs"`
		Timestamp     time.Time `json:"timestamp" sql:"default:now()"`
	}
)
