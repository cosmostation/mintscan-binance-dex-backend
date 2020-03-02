package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"
	cmtypes "github.com/binance-chain/go-sdk/common/types"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"

	amino "github.com/tendermint/go-amino"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data
type Client struct {
	acceleratedNode string
	apiClient       *resty.Client
	cdc             *amino.Codec
	rpcClient       rpc.Client
}

// NewClient returns Client
func NewClient(rpcNode, acceleratedNode, apiServerEndpoint string, networkType cmtypes.ChainNetwork) Client {
	rpcClient := rpc.NewRPCClient(rpcNode, networkType)

	restyClient := resty.New().
		SetHostURL(apiServerEndpoint).
		SetTimeout(time.Duration(5 * time.Second))

	return Client{acceleratedNode, restyClient, codec.Codec, rpcClient}
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

// ValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c Client) ValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(&height)
}

// Validators returns validators detail information in Tendemrint validators in active chain
// An error is returns if the query fails.
func (c Client) Validators() ([]*types.Validator, error) {
	resp, err := c.apiClient.R().Get("/api/v1/stake/validators")
	if err != nil {
		return nil, err
	}

	var vals []*types.Validator

	err = json.Unmarshal(resp.Body(), &vals)
	if err != nil {
		return nil, err
	}

	return vals, nil
}

// Tokens returns information about existing tokens in active chain
func (c Client) Tokens(limit int, offset int) ([]*types.Token, error) {
	resp, err := c.apiClient.R().Get("/api/v1/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	var tokens []*types.Token
	err = json.Unmarshal(resp.Body(), &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
