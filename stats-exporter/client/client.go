package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"
	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/types"

	amino "github.com/tendermint/go-amino"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data
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

// AssetInfoList returns asset info list in active chain
// An error returns if the query fails.
func (c Client) AssetInfoList(page int, rows int) (types.AssetInfo, error) {
	resp, err := c.explorerClient.R().Get("/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return types.AssetInfo{}, err
	}

	var assets types.AssetInfo
	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		return types.AssetInfo{}, err
	}

	return assets, nil
}
