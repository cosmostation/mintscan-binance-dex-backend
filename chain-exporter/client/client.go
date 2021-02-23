package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/binance-chain/go-sdk/client/rpc"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"

	amino "github.com/tendermint/go-amino"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	resty "github.com/go-resty/resty/v2"
)

// Client wraps around both Tendermint RPC client and
// Cosmos SDK LCD REST client that enables to query necessary data.
type Client struct {
	acceleratedClient *resty.Client
	apiClient         *resty.Client
	cdc               *amino.Codec
	explorerClient    *resty.Client
	rpcClient         rpc.Client
}

var (
	controler = make(chan struct{}, 40)
	wg        = new(sync.WaitGroup)
)

// NewClient creates a new Client with the given config.
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

// GetBlock queries for a block by height. An error is returned if the query fails.
func (c Client) GetBlock(height int64) (*tmctypes.ResultBlock, error) {
	return c.rpcClient.Block(&height)
}

// GetLatestBlockHeight returns the latest block height on the active chain.
func (c Client) GetLatestBlockHeight() (int64, error) {
	status, err := c.rpcClient.Status()
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight

	return height, nil
}

// GetTxs queries for all the transactions in a block height.
// It uses `Tx` RPC method to query for the transaction.
func (c Client) GetTxs(block *tmctypes.ResultBlock) ([]*rpc.ResultTx, error) {
	txs := make([]*rpc.ResultTx, len(block.Block.Txs), len(block.Block.Txs))
	var err error
	retryFlag := false

	for i, tmTx := range block.Block.Txs {
		hash := tmTx.Hash()
		controler <- struct{}{}
		wg.Add(1)
		go func(i int, hash []byte) {
			defer func() {
				<-controler
				wg.Done()
			}()

			txs[i], err = c.rpcClient.Tx(hash, true)
			if err != nil {
				retryFlag = true
				fmt.Println(hash)
				return
			}
		}(i, hash)
	}
	wg.Wait()

	if retryFlag {
		return nil, fmt.Errorf("can not get all of txs, retry get tx in block height = %d", block.Block.Height)
	}

	// tx, err := c.rpcClient.Tx(hash, true)
	// if err != nil {
	// 	return nil, err
	// }

	// txs[i] = tx

	return txs, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c Client) GetValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(&height)
}

// GetValidators returns validators detail information in Tendemrint validators in active chain
// An error returns if the query fails.
func (c Client) GetValidators() ([]*types.Validator, error) {
	resp, err := c.apiClient.R().Get("/stake/validators")
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
	resp, err := c.apiClient.R().Get("/tokens?limit=" + strconv.Itoa(limit) + "&offset=" + strconv.Itoa(offset))
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
