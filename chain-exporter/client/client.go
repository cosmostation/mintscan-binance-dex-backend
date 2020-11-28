package client

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/types"
)

type Client struct {
	cdc            *codec.LegacyAmino
	rpcClient      rpcclient.Client
	exchangeClient *resty.Client
}

// NewClient creates a new Client with the given config.
func NewClient(cfg config.NodeConfig) *Client {

	exchangeClient := resty.New().
		SetHostURL(cfg.ExchangeAPIEndpoint).
		SetTimeout(time.Duration(10 * time.Second))

	rpcClient, err := rpchttp.NewWithTimeout(cfg.RPCNode, "/websocket", 10)
	if err != nil {
		panic("failed to init rpcClient: " + err.Error())
	}

	return &Client{
		legacy.Cdc,
		rpcClient,
		exchangeClient,
	}
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c Client) GetBlock(height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(context.Background(), &height)
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func (c Client) GetLatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status(context.Background())
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

// GetTxs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction.
func (c Client) GetTxs(block *tmctypes.ResultBlock) ([]*ctypes.ResultTx, error) {
	txs := make([]*ctypes.ResultTx, len(block.Block.Txs), len(block.Block.Txs))

	for i, tmTx := range block.Block.Txs {
		tx, err := c.rpcClient.Tx(context.Background(), tmTx.Hash(), true)
		if err != nil {
			return nil, err
		}

		txs[i] = tx
	}

	return txs, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c Client) GetValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(context.Background(), &height, nil, nil)
}

// GetValidators returns validators detail information in Tendemrint validators in active chain
// An error returns if the query fails.
func (c Client) GetValidators() ([]*types.Validator, error) {
	resp, err := c.exchangeClient.R().Get("/stake/validators")
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

// GetTokens returns information about existing tokens in active chain.
func (c Client) GetTokens(limit int, offset int) ([]*types.Token, error) {
	resp, err := c.exchangeClient.R().Get("/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
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
