package models

import (
	"encoding/json"
)

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
