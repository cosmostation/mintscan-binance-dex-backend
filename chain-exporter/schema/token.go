package schema

// Token represents token information
type Token struct {
	ID             int32  `json:"id" sql:",pk"`
	Name           string `json:"name" sql:",notnull"`
	Symbol         string `json:"symbol" sql:",notnull, unique"`
	OriginalSymbol string `json:"original_symbol" sql:",notnull"`
	TotalSupply    string `json:"total_supply" sql:",notnull"`
	Owner          string `json:"owner" sql:",notnull"`
	Mintable       bool   `json:"mintable" sql:",notnull"`
}
