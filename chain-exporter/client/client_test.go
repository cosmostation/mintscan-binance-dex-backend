package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"

	resty "github.com/go-resty/resty/v2"
)

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

func TestRequestTokens2(t *testing.T) {
	cfg := config.ParseConfig()

	client := resty.New().SetHostURL(cfg.Node.APIServerEndpoint)

	limit := int(100)
	offset := int(0)

	result := make([]*types.Token, 0)

	resp, err := client.R().Get("/api/v1/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		t.Errorf("failed to request api call: %v\n", err)
	}

	var tks []*types.Token
	err = json.Unmarshal(resp.Body(), &tks)
	if err != nil {
		t.Errorf("failed to unmarshal tokens: %v\n", err)
	}

	for {
		fmt.Println("offset: ", offset)

		offset += limit

		fmt.Println("offset: ", offset)

		if len(tks) < limit {
			break
		}

		resp, err := client.R().Get("/api/v1/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
		if err != nil {
			t.Errorf("failed to request api call: %v\n", err)
		}

		var tks []*types.Token
		err = json.Unmarshal(resp.Body(), &tks)
		if err != nil {
			t.Errorf("failed to unmarshal tokens: %v\n", err)
		}

		for _, tk := range tks {
			result = append(result, tk)
		}
	}

	fmt.Println("length: ", len(result))
}
