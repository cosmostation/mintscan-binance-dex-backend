package models

// MiniTokens defines the structure for BEP8 mini token information.
type MiniTokens struct {
	Name           string `json:"name"`
	OriginalSymbol string `json:"original_symbol"`
	Symbol         string `json:"symbol"`
	Owner          string `json:"owner"`
	TokenURI       string `json:"token_uri"`
	TokenType      int    `json:"token_type"`
	TotalSupply    string `json:"total_supply"`
	Mintable       bool   `json:"mintable"`
}
