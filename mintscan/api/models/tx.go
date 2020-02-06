package models

import (
	"encoding/json"
	"time"
)

// Txs is transaction data for result block
type Txs struct {
	Height    int64     `json:"height"`
	Result    bool      `json:"result"`
	TxHash    string    `json:"tx_hash"`
	Messages  []Message `json:"messages"`
	Memo      string    `json:"memo"`
	Code      uint32    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

// Message describes tx meesage
type Message struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

// Message describes tx signature
type Signature struct {
	Pubkey        string `json:"pubkey"`
	Address       string `json:"address"`
	Sequence      string `json:"sequence"`
	Signature     string `json:"signature"`
	AccountNumber string `json:"account_number"`
}
