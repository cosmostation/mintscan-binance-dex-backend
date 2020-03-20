package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"
	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/models"

	amino "github.com/tendermint/go-amino"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC and
// Cosmos SDK REST clients that enable to query necessary data
type Client struct {
	acceleratedNode string
	apiClient       *resty.Client
	cdc             *amino.Codec
	explorerClient  *resty.Client
	rpcClient       rpc.Client
}

// NewClient returns Client
func NewClient(rpcNode, acceleratedNode, apiServerEndpoint string, explorerServerEndpoint string,
	networkType cmtypes.ChainNetwork) Client {

	restyClient := resty.New().
		SetHostURL(apiServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	explorerClient := resty.New().
		SetHostURL(explorerServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	rpcClient := rpc.NewRPCClient(rpcNode, networkType)

	return Client{
		acceleratedNode,
		restyClient,
		codec.Codec,
		explorerClient,
		rpcClient,
	}
}

// Assets fetches asset list information from an explorer API
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

// Asset fetches particular asset information from an explorer API
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
