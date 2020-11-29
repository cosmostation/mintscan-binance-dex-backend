package types

// Signature defines the structure for transaction signature
type Signature struct {
	Address   string `json:"address,omitempty"`
	Pubkey    string `json:"pubkey,omitempty"`
	Sequence  uint64 `json:"sequence,omitempty"`
	Signature string `json:"signature,omitempty"`
}
