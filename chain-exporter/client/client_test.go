package client

import (
	"encoding/json"
	"testing"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"

	resty "github.com/go-resty/resty/v2"
)

// TestRequestTokens checks to see if token request works
func TestRequestTokens(t *testing.T) {
	cfg := config.ParseConfig()

	client := resty.New().SetHostURL(cfg.Node.APIServerEndpoint)

	resp, err := client.R().Get("/api/v1/tokens?limit=2")
	if err != nil {
		t.Errorf("failed to request api call: %v\n", err)
	}

	var tks []types.Token
	err = json.Unmarshal(resp.Body(), &tks)
	if err != nil {
		t.Errorf("failed to unmarshal tokens: %v\n", err)
	}

	for _, tk := range tks {
		t.Log(tk.Name)
		t.Log(tk.OriginalSymbol)
		t.Log(tk.Owner)
		t.Log(tk.Mintable)
	}
}
