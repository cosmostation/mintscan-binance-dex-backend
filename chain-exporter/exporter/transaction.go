package exporter

import (
	"encoding/base64"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/schema"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// getTxs parses transactions in a block and return transactions.
func (ex *Exporter) getTxs(block *tmctypes.ResultBlock) (transactions []*schema.Transaction, err error) {
	txs, err := ex.client.GetTxs(block)
	if err != nil {
		return []*schema.Transaction{}, err
	}

	if len(txs) <= 0 {
		return []*schema.Transaction{}, nil
	}

	for _, tx := range txs {
		var stdTx txtypes.StdTx
		err = ex.cdc.UnmarshalBinaryLengthPrefixed([]byte(tx.Tx), &stdTx)
		if err != nil {
			return []*schema.Transaction{}, err
		}

		msgsBz, err := ex.cdc.MarshalJSON(stdTx.GetMsgs())
		if err != nil {
			return []*schema.Transaction{}, err
		}

		sigs := make([]types.Signature, len(stdTx.Signatures), len(stdTx.Signatures))

		for i, sig := range stdTx.Signatures {
			consPubKey := sdk.GetConsAddress(sig.PubKey).String()
			if err != nil {
				return []*schema.Transaction{}, err
			}

			sigs[i] = types.Signature{
				Address: sig.Address().String(), // hex string
				// AccountNumber: sig.AccountNumber,
				Pubkey: consPubKey,
				// Sequence:      sig.Sequence,
				Signature: base64.StdEncoding.EncodeToString(sig.Signature), // encode base64
			}
		}

		sigsBz, err := ex.cdc.MarshalJSON(sigs)
		if err != nil {
			return []*schema.Transaction{}, err
		}

		t := &schema.Transaction{
			Height:     tx.Height,
			TxHash:     tx.Hash.String(),
			Code:       tx.TxResult.Code,
			Messages:   string(msgsBz),
			Signatures: string(sigsBz),
			Memo:       stdTx.Memo,
			GasWanted:  tx.TxResult.GasWanted,
			GasUsed:    tx.TxResult.GasUsed,
			Timestamp:  block.Block.Time,
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
