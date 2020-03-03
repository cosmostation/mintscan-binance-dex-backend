package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"

	cmtypes "github.com/binance-chain/go-sdk/common/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	amino "github.com/tendermint/go-amino"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data
type Client struct {
	acceleratedNode string
	apiClient       *resty.Client
	cdc             *amino.Codec
	coinGeckoClient *resty.Client
	rpcClient       rpc.Client
}

// NewClient returns Client
func NewClient(rpcNode, acceleratedNode, apiServerEndpoint string, coinGeckoEndpoint string, networkType cmtypes.ChainNetwork) Client {
	rpcClient := rpc.NewRPCClient(rpcNode, networkType)

	apiClient := resty.New().
		SetHostURL(apiServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	coinGeckoClient := resty.New().
		SetHostURL(coinGeckoEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	return Client{acceleratedNode, apiClient, codec.Codec, coinGeckoClient, rpcClient}
}

// Status returns status info on the active chain
func (c Client) Status() (*ctypes.ResultStatus, error) {
	return c.rpcClient.Status()
}

// Block queries for a block by height. An error is returned if the query fails.
func (c Client) Block(height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(&height)
}

// LatestBlockHeight returns the latest block height on the active chain
func (c Client) LatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

// Tokens returns information about existing tokens in active chain
func (c Client) Tokens(limit int, offset int) ([]*models.Token, error) {
	resp, err := c.apiClient.R().Get("/api/v1/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	var tokens []*models.Token
	err = json.Unmarshal(resp.Body(), &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

// ValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c Client) ValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(&height)
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

// BinanceCoinMarketData fetches BNB market data from CoinGecko API
func (c Client) BinanceCoinMarketData() (models.Market, error) {
	resp, err := c.coinGeckoClient.R().Get("/api/v3/coins/binancecoin")
	if err != nil {
		return models.Market{}, err
	}

	var coingeckoMarket models.CoinGeckoMarket

	err = json.Unmarshal(resp.Body(), &coingeckoMarket)
	if err != nil {
		return models.Market{}, err
	}

	return models.Market{}, nil
}
