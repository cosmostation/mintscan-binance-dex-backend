package exporter

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"
)

// 배열을 만들고
// 요청 후 10개 면 10개 가져온 후 배열에 담고 offset 10개 올린 후 다시 요청

// getTokens
func (ex *Exporter) getTokens() ([]*schema.Token, error) {
	limit := 100
	offset := 0

	tokens := make([]*types.Token, 0)

	tks, err := ex.client.Tokens(limit, offset)
	if err != nil {
		return nil, err
	}

	for _, tk := range tks {
		tokens = append(tokens, tk)
	}

	for {
		// Increase offset by limit to request
		offset += limit

		if len(tks) < offset {
			break
		}

		tks, err := ex.client.Tokens(limit, offset)
		if err != nil {
			return nil, err
		}

		for _, tk := range tks {
			tokens = append(tokens, tk)
		}
	}

	return tokens, nil
}
