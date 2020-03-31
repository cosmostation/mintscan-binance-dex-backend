package types

// Token defines the structure for token information
type Token struct {
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	OriginalSymbol string `json:"original_symbol"`
	TotalSupply    string `json:"total_supply"`
	Owner          string `json:"owner"`
	Mintable       bool   `json:"mintable"`
}
