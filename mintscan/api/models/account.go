package models

import "encoding/json"

// Account represents account information
type Account struct {
	Address       string          `json:"address"`
	PublicKey     json.RawMessage `json:"public_key"`
	AccountNumber int64           `json:"account_number"`
	Sequence      int64           `json:"sequence"`
	Flags         uint64          `json:"flags"`
	Balances      []struct {
		Symbol string `json:"symbol"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
		Frozen string `json:"frozen"`
	} `json:"balances"`
}
