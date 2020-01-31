package client

import (
	"github.com/binance-chain/go-sdk/client/rpc"
	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/codec"

	amino "github.com/tendermint/go-amino"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
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

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

// Txs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction
func (c Client) Txs(block *tmctypes.ResultBlock) ([]*rpc.ResultTx, error) {
	txs := make([]*rpc.ResultTx, len(block.Block.Txs), len(block.Block.Txs))

	for i, tmTx := range block.Block.Txs {
		tx, err := c.rpcClient.Tx(tmTx.Hash(), true)
		if err != nil {
			return nil, err
		}

		txs[i] = tx
	}

	return txs, nil
}

// Validators returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c Client) Validators(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(&height)
}
