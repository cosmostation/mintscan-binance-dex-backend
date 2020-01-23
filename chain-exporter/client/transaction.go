package client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Tx queries for a transaction from the REST client and decodes it into a sdk.TxResponse
func (c Client) Tx(hash string) (sdk.TxResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/txs/%s", c.lcdClient, hash))
	if err != nil {
		return sdk.TxResponse{}, err
	}

	defer resp.Body.Close()

	bz, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	var tx sdk.TxResponse
	if err := c.cdc.UnmarshalJSON(bz, &tx); err != nil {
		return sdk.TxResponse{}, err
	}

	return tx, nil
}

// Txs queries for all the transactions in a block
func (c Client) Txs(block *tmctypes.ResultBlock) ([]sdk.TxResponse, error) {
	txResponses := make([]sdk.TxResponse, len(block.Block.Txs), len(block.Block.Txs))

	for i, tmTx := range block.Block.Txs {
		txResponse, err := c.Tx(fmt.Sprintf("%X", tmTx.Hash()))
		if err != nil {
			return nil, err
		}

		txResponses[i] = txResponse
	}

	return txResponses, nil
}
