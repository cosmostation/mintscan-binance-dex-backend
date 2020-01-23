package client

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/codec"

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
func NewClient(rpcNode, lcdEndpoint string) Client {
	return Client{rpc.NewRPCClient(rpcNode, cmtypes.TestNetwork), lcdEndpoint, codec.Codec}
}
