package models

import "time"

// Paging wraps required params for handling pagination
type Paging struct {
	Total  int   `json:"total"` // total number of txs saved in database
	Before int32 `json:"before"`
	After  int32 `json:"after"`
}

type (
	// ResultBlocks is a block result response
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

type (
	// ResultTxs is transaction result response
	ResultTxs struct {
		Paging Paging   `json:"paging"`
		Data   []TxData `json:"data"`
	}

	// TxData wraps tx data
	TxData struct {
		ID         int32       `json:"id"`
		Height     int64       `json:"height"`
		Result     bool        `json:"result"`
		TxHash     string      `json:"tx_hash"`
		Messages   []Message   `json:"messages"`
		Signatures []Signature `json:"signatures"`
		Memo       string      `json:"memo"`
		Code       uint32      `json:"code"`
		Timestamp  time.Time   `json:"timestamp"`
	}
)
