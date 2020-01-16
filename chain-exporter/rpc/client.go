package rpc

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	"github.com/binance-chain/go-sdk/common/types"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
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

// LatestBlockHeight returns the latest block height on the active chain
func (c Client) LatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

// Block queries for a block by height. An error is returned if the query fails.
func (c Client) Block(height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(&height)
}
