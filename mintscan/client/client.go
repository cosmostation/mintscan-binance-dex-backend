package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	amino "github.com/tendermint/go-amino"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps for both Tendermint RPC and other API clients that
// are needed for this project
type Client struct {
	acceleratedClient *resty.Client
	apiClient         *resty.Client
	cdc               *amino.Codec
	coinGeckoClient   *resty.Client
	explorerClient    *resty.Client
	rpcClient         rpc.Client
}

// NewClient creates a new client with the given config
func NewClient(cfg config.NodeConfig, marketCfg config.MarketConfig) *Client {

	acceleratedClient := resty.New().
		SetHostURL(cfg.AcceleratedNode).
		SetTimeout(time.Duration(5 * time.Second))

	apiClient := resty.New().
		SetHostURL(cfg.APIServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	coinGeckoClient := resty.New().
		SetHostURL(marketCfg.CoinGeckoEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	explorerClient := resty.New().
		SetHostURL(cfg.ExplorerServerEndpoint).
		SetTimeout(time.Duration(50 * time.Second))

	rpcClient := rpc.NewRPCClient(cfg.RPCNode, cfg.NetworkType)

	return &Client{
		acceleratedClient,
		apiClient,
		codec.Codec,
		coinGeckoClient,
		explorerClient,
		rpcClient,
	}
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
	resp, err := c.apiClient.R().Get("/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
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
	resp, err := c.apiClient.R().Get("/stake/validators")
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

// CoinMarketData returns current market data from CoinGecko API based upon params
func (c Client) CoinMarketData(id string) (models.CoinGeckoMarket, error) {
	queryStr := "/coins/" + id + "?localization=false&tickers=false&community_data=false&developer_data=false&sparkline=false"

	resp, err := c.coinGeckoClient.R().Get(queryStr)
	if err != nil {
		return models.CoinGeckoMarket{}, err
	}

	if resp.IsError() {
		return models.CoinGeckoMarket{}, fmt.Errorf("failed to respond: %s", err)
	}

	var data models.CoinGeckoMarket
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return models.CoinGeckoMarket{}, err
	}

	return data, nil
}

// CoinMarketChartData returns current market chart data from CoinGecko API based upon params
func (c Client) CoinMarketChartData(id string, from string, to string) (models.CoinGeckoMarketChart, error) {
	queryStr := "/coins/" + id + "/market_chart/range?id=" + id + "&vs_currency=usd&from=" + from + "&to=" + to

	resp, err := c.coinGeckoClient.R().Get(queryStr)
	if err != nil {
		return models.CoinGeckoMarketChart{}, err
	}

	if resp.IsError() {
		return models.CoinGeckoMarketChart{}, fmt.Errorf("failed to respond: %s", err)
	}

	var data models.CoinGeckoMarketChart
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return models.CoinGeckoMarketChart{}, err
	}

	return data, nil
}

// Asset returns particular asset information given an asset name
func (c Client) Asset(assetName string) (models.Asset, error) {
	resp, err := c.explorerClient.R().Get("/asset?asset=" + assetName)
	if err != nil {
		return models.Asset{}, err
	}

	var asset models.Asset
	err = json.Unmarshal(resp.Body(), &asset)
	if err != nil {
		return models.Asset{}, err
	}

	return asset, nil
}

// Assets returns information of all assets existing in an active chain
func (c Client) Assets(page int, rows int) (models.AssetInfo, error) {
	queryStr := "/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows)
	resp, err := c.explorerClient.R().Get(queryStr)
	if err != nil {
		return models.AssetInfo{}, err
	}

	var assets models.AssetInfo
	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		return models.AssetInfo{}, err
	}

	return assets, nil
}

// AssetHolders returns all asset holders information based upon params
func (c Client) AssetHolders(asset string, page int, rows int) (models.AssetHolders, error) {
	queryStr := "/asset-holders?asset=" + asset + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows)
	resp, err := c.explorerClient.R().Get(queryStr)
	if err != nil {
		return models.AssetHolders{}, err
	}

	var assetHolders models.AssetHolders
	err = json.Unmarshal(resp.Body(), &assetHolders)
	if err != nil {
		return models.AssetHolders{}, err
	}

	return assetHolders, nil
}

// AssetTxs returns asset transactions given an asset name based upon params
func (c Client) AssetTxs(txAsset string, page int, rows int) (models.AssetTxs, error) {
	queryStr := "/txs?txAsset=" + txAsset + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows)
	resp, err := c.explorerClient.R().Get(queryStr)
	if err != nil {
		return models.AssetTxs{}, err
	}

	var assetTxs models.AssetTxs
	err = json.Unmarshal(resp.Body(), &assetTxs)
	if err != nil {
		return models.AssetTxs{}, err
	}

	return assetTxs, nil
}

// Account returns account information given an account address
func (c Client) Account(address string) (models.Account, error) {
	resp, err := c.apiClient.R().Get("/account/" + address)
	if err != nil {
		return models.Account{}, err
	}

	var account models.Account
	err = json.Unmarshal(resp.Body(), &account)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// AccountTxs retuns tranctions involving in an account based upon params
func (c Client) AccountTxs(address string, page int, rows int) (models.AccountTxs, error) {
	queryStr := "/txs?address=" + address + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows)
	resp, err := c.explorerClient.R().Get(queryStr)
	if err != nil {
		return models.AccountTxs{}, err
	}

	var acctTxs models.AccountTxs
	err = json.Unmarshal(resp.Body(), &acctTxs)
	if err != nil {
		return models.AccountTxs{}, err
	}

	return acctTxs, nil
}

// Order returns order information given an order id
func (c Client) Order(id string) (models.Order, error) {
	resp, err := c.acceleratedClient.R().Get("/orders/" + id)
	if err != nil {
		return models.Order{}, err
	}

	var order models.Order
	err = json.Unmarshal(resp.Body(), &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// TxMsgFees returns fees for different transaciton message types
func (c Client) TxMsgFees() ([]*models.TxMsgFee, error) {
	resp, err := c.acceleratedClient.R().Get("/fees")
	if err != nil {
		return []*models.TxMsgFee{}, err
	}

	var fees []*models.TxMsgFee
	err = json.Unmarshal(resp.Body(), &fees)
	if err != nil {
		return []*models.TxMsgFee{}, err
	}

	return fees, nil
}
