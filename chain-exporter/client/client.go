package client

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data
type Client struct {
	rpcClient rpc.Client
	lcdClient string
}

// NewClient returns Client
func NewClient(rpcNode, lcdEndpoint string) Client {
	return Client{rpc.NewRPCClient(rpcNode, types.TestNetwork), lcdEndpoint}
}
