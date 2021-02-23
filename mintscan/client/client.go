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

// --------------------
// RPC APIs
// --------------------

// GetStatus returns status info on the active chain.
func (c *Client) GetStatus() (*ctypes.ResultStatus, error) {
	return c.rpcClient.Status()
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c *Client) GetBlock(height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(&height)
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func (c *Client) GetLatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c *Client) GetValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(&height)
}

// --------------------
// REST SERVER APIs
// --------------------

// GetTokens returns information about existing tokens in active chain
func (c *Client) GetTokens(limit int, offset int) (tokens []models.Token, err error) {
	resp, err := c.apiClient.R().Get("/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		return []models.Token{}, err
	}

	if resp.IsError() {
		return []models.Token{}, fmt.Errorf("failed to request tokens: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &tokens)
	if err != nil {
		return []models.Token{}, err
	}

	return tokens, nil
}

// GetValidators returns validators detail information in Tendemrint validators in active chain
// An error is returns if the query fails.
func (c *Client) GetValidators() (validators []models.Validator, err error) {
	resp, err := c.apiClient.R().Get("/stake/validators")
	if err != nil {
		return []models.Validator{}, err
	}

	if resp.IsError() {
		return []models.Validator{}, fmt.Errorf("failed to request validators: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &validators)
	if err != nil {
		return []models.Validator{}, err
	}

	return validators, nil
}

// GetCoinMarketData returns market data from CoinGecko API based upon params.
func (c *Client) GetCoinMarketData(id string) (data models.CoinGeckoMarket, err error) {
	resp, err := c.coinGeckoClient.R().Get("/coins/" + id + "?localization=false&tickers=false&community_data=false&developer_data=false&sparkline=false")
	if err != nil {
		return models.CoinGeckoMarket{}, err
	}

	if resp.IsError() {
		return models.CoinGeckoMarket{}, fmt.Errorf("failed to request market data: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return models.CoinGeckoMarket{}, err
	}

	return data, nil
}

// GetCoinMarketChartData returns current market chart data from CoinGecko API based upon params.
func (c *Client) GetCoinMarketChartData(id string, from string, to string) (data models.CoinGeckoMarketChart, err error) {
	queryStr := "/coins/" + id + "/market_chart/range?id=" + id + "&vs_currency=usd&from=" + from + "&to=" + to

	resp, err := c.coinGeckoClient.R().Get(queryStr)
	if err != nil {
		return models.CoinGeckoMarketChart{}, err
	}

	if resp.IsError() {
		return models.CoinGeckoMarketChart{}, fmt.Errorf("failed to request chart data: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return models.CoinGeckoMarketChart{}, err
	}

	return data, nil
}

// GetAsset returns particular asset information given an asset name.
func (c *Client) GetAsset(assetName string) (asset models.Asset, err error) {
	resp, err := c.explorerClient.R().Get("/asset?asset=" + assetName)
	if err != nil {
		return models.Asset{}, err
	}

	if resp.IsError() {
		return models.Asset{}, fmt.Errorf("failed to request asset: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &asset)
	if err != nil {
		return models.Asset{}, err
	}

	return asset, nil
}

// GetAssets returns information of all assets existing in an active chain.
func (c *Client) GetAssets(page int, rows int) (assets models.AssetInfo, err error) {
	resp, err := c.explorerClient.R().Get("/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.AssetInfo{}, err
	}

	if resp.IsError() {
		return models.AssetInfo{}, fmt.Errorf("failed to request assets: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		return models.AssetInfo{}, err
	}

	return assets, nil
}

// GetAssetHolders returns all asset holders information based upon params.
func (c *Client) GetAssetHolders(asset string, page int, rows int) (holders models.AssetHolders, err error) {
	resp, err := c.explorerClient.R().Get("/asset-holders?asset=" + asset + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.AssetHolders{}, err
	}

	if resp.IsError() {
		return models.AssetHolders{}, fmt.Errorf("failed to request asset holders: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &holders)
	if err != nil {
		return models.AssetHolders{}, err
	}

	return holders, nil
}

// GetAssetTxs returns asset transactions given an asset name based upon params.
func (c *Client) GetAssetTxs(txAsset string, page int, rows int) (txs models.AssetTxs, err error) {
	resp, err := c.explorerClient.R().Get("/txs?txAsset=" + txAsset + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.AssetTxs{}, err
	}

	if resp.IsError() {
		return models.AssetTxs{}, fmt.Errorf("failed to request asset transactions: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &txs)
	if err != nil {
		return models.AssetTxs{}, err
	}

	return txs, nil
}

// GetMiniTokens returns a list of available mini tokens.
func (c *Client) GetMiniTokens(page int, rows int) (assets models.AssetInfo, err error) {
	resp, err := c.explorerClient.R().Get("/mini-token/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.AssetInfo{}, err
	}

	if resp.IsError() {
		return models.AssetInfo{}, fmt.Errorf("failed to request bep8 tokens: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		return models.AssetInfo{}, err
	}

	return assets, nil
}

// GetAccount returns account information given an account address.
func (c *Client) GetAccount(address string) (account models.Account, err error) {
	resp, err := c.apiClient.R().Get("/account/" + address)
	if err != nil {
		return models.Account{}, err
	}

	if resp.IsError() {
		return models.Account{}, fmt.Errorf("failed to request account information: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &account)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}

// GetAccountTxs retuns tranctions involving in an account based upon params.
func (c *Client) GetAccountTxs(address string, page int, rows int) (txs models.AccountTxs, err error) {
	resp, err := c.explorerClient.R().Get("/txs?address=" + address + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.AccountTxs{}, err
	}

	if resp.IsError() {
		return models.AccountTxs{}, fmt.Errorf("failed to request account transactinos: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &txs)
	if err != nil {
		return models.AccountTxs{}, err
	}

	return txs, nil
}

// GetOrder returns order information with given order id.
func (c *Client) GetOrder(id string) (order models.Order, err error) {
	var resp *resty.Response

	resp, err = c.acceleratedClient.R().Get("/orders/" + id)
	if err != nil {
		return models.Order{}, err
	}

	if resp.IsError() {
		return models.Order{}, fmt.Errorf("failed to request bep2 token order information: %d", resp.StatusCode())
	}

	if resp.String() == "" {
		resp, err = c.acceleratedClient.R().Get("/mini/orders/" + id)
		if err != nil {
			return models.Order{}, err
		}

		if resp.IsError() {
			return models.Order{}, fmt.Errorf("failed to request bep8 token order information: %d", resp.StatusCode())
		}
	}

	err = json.Unmarshal(resp.Body(), &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// GetTxMsgFees returns fees for different transaciton message types.
func (c *Client) GetTxMsgFees() (fees []*models.TxMsgFee, err error) {
	resp, err := c.acceleratedClient.R().Get("/fees")
	if err != nil {
		return []*models.TxMsgFee{}, err
	}

	if resp.IsError() {
		return []*models.TxMsgFee{}, fmt.Errorf("failed to request tx msg fees: %d", resp.StatusCode())
	}

	err = json.Unmarshal(resp.Body(), &fees)
	if err != nil {
		return []*models.TxMsgFee{}, err
	}

	return fees, nil
}
