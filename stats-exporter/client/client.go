package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"

	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/models"

	amino "github.com/tendermint/go-amino"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC and
// Cosmos SDK REST clients that enable to query necessary data
type Client struct {
	acceleratedClient *resty.Client
	apiClient         *resty.Client
	cdc               *amino.Codec
	explorerClient    *resty.Client
	rpcClient         rpc.Client
}

// NewClient creates a new client with the given config
func NewClient(cfg config.NodeConfig) *Client {

	acceleratedClient := resty.New().
		SetHostURL(cfg.AcceleratedNode).
		SetTimeout(time.Duration(5 * time.Second))

	apiClient := resty.New().
		SetHostURL(cfg.APIServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	explorerClient := resty.New().
		SetHostURL(cfg.ExplorerServerEndpoint).
		SetTimeout(time.Duration(30 * time.Second))

	rpcClient := rpc.NewRPCClient(cfg.RPCNode, cfg.NetworkType)

	return &Client{
		acceleratedClient,
		apiClient,
		codec.Codec,
		explorerClient,
		rpcClient,
	}
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

// Assets returns information of all assets existing in an active chain based upon params
func (c Client) Assets(page int, rows int) (models.Assets, error) {
	resp, err := c.explorerClient.R().Get("/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.Assets{}, err
	}

	var assets models.Assets
	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		return models.Assets{}, err
	}

	return assets, nil
}
