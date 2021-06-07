package client

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/types"
	chainClient "github.com/InjectiveLabs/sdk-go/chain/client"
)

type Client struct {
	ctx          client.Context
	cosmosClient chainClient.CosmosClient
	rpcClient    rpcclient.Client
}

// NewClient creates a new Client with the given config.
func NewClient(cfg config.NodeConfig) *Client {
	rpcClient, err := rpchttp.NewWithTimeout(cfg.RPCNode, "/websocket", 10)
	if err != nil {
		log.WithError(err).Fatalln("failed to init rpcClient")
	}

	ctx, err := chainClient.NewClientContext(cfg.ChainID, "", nil)
	if err != nil {
		log.WithError(err).Fatalln("failed to init cosmos client context")
	}

	cosmosClient, err := chainClient.NewCosmosClient(ctx, cfg.GRPCNode)
	if err != nil {
		log.WithError(err).Fatalln("failed to connect to the gRPC")
	}

	return &Client{
		ctx:          ctx,
		cosmosClient: cosmosClient,
		rpcClient:    rpcClient,
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
	txs := make([]*ctypes.ResultTx, 0, len(block.Block.Txs))

	for _, tmTx := range block.Block.Txs {
		tx, err := c.rpcClient.Tx(context.Background(), tmTx.Hash(), true)
		if err != nil {
			if strings.HasSuffix(err.Error(), "not found") {
				log.WithError(err).Errorln("failed to get Tx by hash")
				continue
			}

			return nil, err
		}

		txs = append(txs, tx)
	}

	return txs, nil
}

// GetValidatorSet returns all the known Tendermint validators for a given block
// height. An error is returned if the query fails.
func (c Client) GetValidatorSet(height int64) (*tmctypes.ResultValidators, error) {
	return c.rpcClient.Validators(context.Background(), &height, nil, nil)
}

func (c Client) TxDecoder() sdk.TxDecoder {
	return c.ctx.TxConfig.TxDecoder()
}

func (c Client) JSONMarshaler() codec.JSONCodec {
	return c.ctx.JSONCodec
}

func (c Client) Marshaler() codectypes.InterfaceRegistry {
	return c.ctx.InterfaceRegistry
}

// GetValidators returns validators detail information in Tendemrint validators in active chain
// An error returns if the query fails.
func (c Client) GetValidators() ([]*types.Validator, error) {
	stakingCli := staking.NewQueryClient(c.cosmosClient.QueryClient())
	resp, err := stakingCli.Validators(context.Background(), &staking.QueryValidatorsRequest{})
	if err != nil {
		err = errors.Wrap(err, "failed to query validators from staking module")
		return nil, err
	}

	vals := make([]*types.Validator, 0, len(resp.Validators))
	for _, val := range resp.Validators {

		v := &types.Validator{
			OperatorAddress: val.OperatorAddress, // string
			Jailed:          val.Jailed,          // bool
			Status:          val.Status.String(), // string
			Tokens:          val.Tokens.String(), // string
			Power:           val.ConsensusPower(sdk.NewInt(1000000000000000000)),
			DelegatorShares: val.DelegatorShares.String(), // string
			Description: types.Description{
				Moniker:  val.Description.Moniker,
				Identity: val.Description.Identity,
				Website:  val.Description.Website,
				Details:  val.Description.Details,
			}, // Description
			UnbondingHeight: val.UnbondingHeight, // int64
			UnbondingTime:   val.UnbondingTime,   // time.Time
			Commission: types.Commission{
				Rate:          val.Commission.Rate.String(),
				MaxRate:       val.Commission.MaxRate.String(),
				MaxChangeRate: val.Commission.MaxChangeRate.String(),
				UpdateTime:    val.Commission.UpdateTime,
			}, // Commission
		}

		var pubKey cryptotypes.PubKey
		if err := c.ctx.InterfaceRegistry.UnpackAny(val.ConsensusPubkey, &pubKey); err != nil {
			err = errors.Wrap(err, "failed to unpack val cons pubkey")
			return nil, err
		}

		v.ConsensusPubKey = sdk.MustBech32ifyAddressBytes(sdk.Bech32PrefixConsPub, pubKey.Bytes())

		vals = append(vals, v)
	}

	return vals, nil
}
