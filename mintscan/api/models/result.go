package models

import "time"

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
