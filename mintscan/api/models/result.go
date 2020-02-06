package models

import "time"

// ResultBlock is a block result response
type ResultBlock struct {
	ID            int       `json:"id"`
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

// ResultTxs is transaction result response
type ResultTxs struct {
	ID         int         `json:"id"`
	Height     int64       `json:"height"`
	Result     bool        `json:"result"`
	TxHash     string      `json:"tx_hash"`
	Messages   []Message   `json:"messages"`
	Signatures []Signature `json:"signatures"`
	Memo       string      `json:"memo"`
	Code       uint32      `json:"code"`
	Timestamp  time.Time   `json:"timestamp"`
}
