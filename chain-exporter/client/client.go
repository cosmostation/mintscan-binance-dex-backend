package client

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/types"
)

type Client struct {
	cdc       *codec.LegacyAmino
	rpcClient rpcclient.Client
	grpcConn  *grpc.ClientConn
}

// NewClient creates a new Client with the given config.
func NewClient(cfg config.NodeConfig) *Client {
	rpcClient, err := rpchttp.NewWithTimeout(cfg.RPCNode, "/websocket", 10)
	if err != nil {
		panic("failed to init rpcClient: " + err.Error())
	}

	grpcConn, err := grpc.Dial(cfg.GRPCNode, grpc.WithInsecure(), grpc.WithContextDialer(dialerFunc))
	if err != nil {
		panic("failed to connect to the gRPC: " + cfg.GRPCNode)
	}

	return &Client{
		cdc:       legacy.Cdc,
		rpcClient: rpcClient,
		grpcConn:  grpcConn,
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
	stakingCli := staking.NewQueryClient(c.grpcConn)
	resp, err := stakingCli.Validators(context.Background(), &staking.QueryValidatorsRequest{})
	if err != nil {
		err = errors.Wrap(err, "failed to query validators from staking module")
		return nil, err
	}

	vals := make([]*types.Validator, 0, len(resp.Validators))
	for _, val := range resp.Validators {
		vals = append(vals, &types.Validator{
			OperatorAddress: val.OperatorAddress, // string
			ConsensusPubKey: val.ConsensusPubkey, // string
			Jailed:          val.Jailed,          // bool
			Status:          val.Status.String(), // string
			Tokens:          val.Tokens.String(), // string
			Power:           val.ConsensusPower(),
			DelegatorShares: val.DelegatorShares.String(), // string
			Description: types.Description{
				Moniker:  val.Description.Moniker,
				Identity: val.Description.Identity,
				Website:  val.Description.Website,
				Details:  val.Description.Details,
			}, // Description
			UnbondingHeight: val.UnbondingHeight,        // int64
			UnbondingTime:   val.UnbondingTime.String(), // string
			Commission: types.Commission{
				Rate:          val.Commission.Rate.String(),
				MaxRate:       val.Commission.MaxRate.String(),
				MaxChangeRate: val.Commission.MaxChangeRate.String(),
				UpdateTime:    val.Commission.UpdateTime,
			}, // Commission
		})
	}

	return vals, nil
}
