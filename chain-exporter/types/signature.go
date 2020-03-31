package types

// Signature defines the structure for transaction signature
type Signature struct {
	Address       string `json:"address,omitempty"`
	AccountNumber int64  `json:"account_number,omitempty"`
	Pubkey        string `json:"pubkey,omitempty"`
	Sequence      int64  `json:"sequence,omitempty"`
	Signature     string `json:"signature,omitempty"`
}
