package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	resty "github.com/go-resty/resty/v2"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
)

// Client wraps for both Tendermint RPC and other API clients that
// are needed for this project
type Client struct {
	rpcClient       rpcclient.Client
	grpcConn        *grpc.ClientConn
	exchangeClient  *resty.Client
	coinGeckoClient *resty.Client
}

// NewClient creates a new client with the given config
func NewClient(cfg config.NodeConfig, marketCfg config.MarketConfig) *Client {
	coinGeckoClient := resty.New().
		SetHostURL(marketCfg.CoinGeckoEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	exchangeClient := resty.New().
		SetHostURL(cfg.ExchangeAPIEndpoint).
		SetTimeout(time.Duration(50 * time.Second))

	rpcClient, err := rpchttp.NewWithTimeout(cfg.RPCNode, "/websocket", 10)
	if err != nil {
		panic("failed to init rpcClient: " + err.Error())
	}

	grpcConn, err := grpc.Dial(cfg.GRPCNode, grpc.WithInsecure(), grpc.WithContextDialer(dialerFunc))
	if err != nil {
		panic("failed to connect to the gRPC: " + cfg.GRPCNode)
	}

	return &Client{
		rpcClient:       rpcClient,
		grpcConn:        grpcConn,
		exchangeClient:  exchangeClient,
		coinGeckoClient: coinGeckoClient,
	}
}

// --------------------
// RPC APIs
// --------------------

// GetStatus returns status info on the active chain.
func (c *Client) GetStatus() (*ctypes.ResultStatus, error) {
	return c.rpcClient.Status(context.Background())
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c *Client) GetBlock(height int64) (*ctypes.ResultBlock, error) {
	return c.rpcClient.Block(context.Background(), &height)
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func (c *Client) GetLatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status(context.Background())
	if err != nil {
		return -1, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c *Client) GetValidatorSet(height int64) (*ctypes.ResultValidators, error) {
	return c.rpcClient.Validators(context.Background(), &height, nil, nil)
}

// -------------
// OFFCHAIN APIs
// -------------

// GetTokens returns information about existing tokens in active chain
func (c *Client) GetTokens(limit int, offset int) (tokens []models.Token, err error) {
	resp, err := c.exchangeClient.R().Get("/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
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

// GetValidators returns validators detail information on Tendermint validators in active chain
// An error is returns if the query fails.
func (c Client) GetValidators() ([]models.Validator, error) {
	stakingCli := staking.NewQueryClient(c.grpcConn)
	resp, err := stakingCli.Validators(context.Background(), &staking.QueryValidatorsRequest{})
	if err != nil {
		err = errors.Wrap(err, "failed to query validators from staking module")
		return []models.Validator{}, err
	}

	vals := make([]models.Validator, 0, len(resp.Validators))
	for _, val := range resp.Validators {
		vals = append(vals, models.Validator{
			OperatorAddress: val.OperatorAddress, // string
			ConsensusPubKey: val.ConsensusPubkey, // string
			Jailed:          val.Jailed,          // bool
			Status:          val.Status.String(), // string
			Tokens:          val.Tokens.String(), // string
			Power:           val.ConsensusPower(),
			DelegatorShares: val.DelegatorShares.String(), // string
			Description: models.Description{
				Moniker:  val.Description.Moniker,
				Identity: val.Description.Identity,
				Website:  val.Description.Website,
				Details:  val.Description.Details,
			}, // Description
			UnbondingHeight: val.UnbondingHeight,        // int64
			UnbondingTime:   val.UnbondingTime.String(), // string
			Commission: models.Commission{
				Rate:          val.Commission.Rate.String(),
				MaxRate:       val.Commission.MaxRate.String(),
				MaxChangeRate: val.Commission.MaxChangeRate.String(),
				UpdateTime:    val.Commission.UpdateTime,
			}, // Commission
		})
	}

	return vals, nil
}

// GetAsset returns particular asset information given an asset name.
func (c *Client) GetAsset(assetName string) (asset models.Asset, err error) {
	resp, err := c.exchangeClient.R().Get("/asset?asset=" + assetName)
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
	resp, err := c.exchangeClient.R().Get("/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
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
	resp, err := c.exchangeClient.R().Get("/asset-holders?asset=" + asset + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
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
	resp, err := c.exchangeClient.R().Get("/txs?txAsset=" + txAsset + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
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
	resp, err := c.exchangeClient.R().Get("/mini-token/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
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
	resp, err := c.exchangeClient.R().Get("/account/" + address)
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
	resp, err := c.exchangeClient.R().Get("/txs?address=" + address + "&page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
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

	resp, err = c.exchangeClient.R().Get("/orders/" + id)
	if err != nil {
		return models.Order{}, err
	}

	if resp.IsError() {
		return models.Order{}, fmt.Errorf("failed to request bep2 token order information: %d", resp.StatusCode())
	}

	if resp.String() == "" {
		resp, err = c.exchangeClient.R().Get("/mini/orders/" + id)
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
	resp, err := c.exchangeClient.R().Get("/fees")
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
