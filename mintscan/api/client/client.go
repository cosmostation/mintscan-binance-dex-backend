package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"

	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"

	amino "github.com/tendermint/go-amino"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data
type Client struct {
	acceleratedNode string
	apiClient       *resty.Client
	cdc             *amino.Codec
	rpcClient       rpc.Client
}

// NewClient returns Client
func NewClient(rpcNode, acceleratedNode, apiServerEndpoint string, networkType cmtypes.ChainNetwork) Client {
	rpcClient := rpc.NewRPCClient(rpcNode, networkType)

	restyClient := resty.New().
		SetHostURL(apiServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	return Client{acceleratedNode, restyClient, codec.Codec, rpcClient}
}

// LatestBlockHeight returns the latest block height on the active chain
func (c Client) LatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

func (c Client) Tokens() ([]*cmtypes.Token, error) {
	resp, err := c.apiClient.R().Get("/api/v1/tokens")
	if err != nil {
		return nil, err
	}

	tokens := make([]*cmtypes.Token, 0)
	err = json.Unmarshal(resp.Body(), &tokens)
	if err != nil {
		return nil, err
	}

	fmt.Println("length: ", len(tokens))
	for _, token := range tokens {
		fmt.Println("Name: ", token.Name)
		fmt.Println("Symbol: ", token.Symbol)
		fmt.Println("OrigSymbol: ", token.OrigSymbol)
		fmt.Println("TotalSupply: ", token.TotalSupply)
		fmt.Println("Owner: ", token.Owner)
		fmt.Println("Mintable: ", token.Mintable)
	}

	return tokens, nil
}

// Validators returns validators detail information in Tendemrint validators in active chain
// An error is returns if the query fails.
func (c Client) Validators() ([]*models.Validator, error) {
	resp, err := c.apiClient.R().Get("/api/v1/stake/validators")
	if err != nil {
		return nil, err
	}

	var vals []*models.Validator

	err = json.Unmarshal(resp.Body(), &vals)
	if err != nil {
		return nil, err
	}

	return vals, nil
}
