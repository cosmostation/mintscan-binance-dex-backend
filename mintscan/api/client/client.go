package client

import (
	"github.com/binance-chain/go-sdk/client/rpc"

	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/codec"

	amino "github.com/tendermint/go-amino"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data
type Client struct {
	rpcClient rpc.Client
	lcdClient string
	cdc       *amino.Codec
}

// NewClient returns Client
func NewClient(rpcNode, lcdEndpoint string, networkType cmtypes.ChainNetwork) Client {
	return Client{rpc.NewRPCClient(rpcNode, networkType), lcdEndpoint, codec.Codec}
}

// LatestBlockHeight returns the latest block height on the active chain
func (c Client) LatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	return status.SyncInfo.LatestBlockHeight, nil
}
