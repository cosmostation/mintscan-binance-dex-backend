package exporter

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
)

// getTokens gets all tokens availble on the active chain
func (ex *Exporter) getTokens() ([]*schema.Token, error) {
	limit := 100
	offset := 0
	lock := false

	tokens := make([]*schema.Token, 0)

	for {
		if lock == true {
			break
		}

		tks, err := ex.client.Tokens(limit, offset)
		if err != nil {
			return nil, err
		}

		for _, tk := range tks {
			ok, err := ex.db.ExistToken(tk.OriginalSymbol)
			if !ok {
				tempTk := &schema.Token{
					Name:           tk.Name,
					Symbol:         tk.Symbol,
					OriginalSymbol: tk.OriginalSymbol,
					TotalSupply:    tk.TotalSupply,
					Owner:          tk.Owner,
					Mintable:       tk.Mintable,
				}

				tokens = append(tokens, tempTk)
			}
			if err != nil {
				return nil, fmt.Errorf("unexpected error when checking token existence: %t", err)
			}
		}

		if len(tks) != limit {
			lock = true
		}

		offset = offset + limit
	}

	return tokens, nil
}
