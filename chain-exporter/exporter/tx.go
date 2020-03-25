package exporter

import (
	"encoding/base64"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	ctypes "github.com/binance-chain/go-sdk/common/types"
	txtypes "github.com/binance-chain/go-sdk/types/tx"
)

// getTxs parses transactions and wrap into Transaction schema struct
func (ex *Exporter) getTxs(block *tmctypes.ResultBlock) ([]*schema.Transaction, error) {
	transactions := make([]*schema.Transaction, 0)

	txs, err := ex.client.Txs(block)
	if err != nil {
		return nil, err
	}

	if len(txs) > 0 {
		for _, tx := range txs {
			var stdTx txtypes.StdTx
			ex.cdc.UnmarshalBinaryLengthPrefixed([]byte(tx.Tx), &stdTx)

			msgsBz, err := ex.cdc.MarshalJSON(stdTx.GetMsgs())
			if err != nil {
				return nil, err
			}

			sigs := make([]types.Signature, len(stdTx.Signatures), len(stdTx.Signatures))

			for i, sig := range stdTx.Signatures {
				consPubKey, err := ctypes.Bech32ifyConsPub(sig.PubKey)
				if err != nil {
					return nil, err
				}

				sigs[i] = types.Signature{
					Address:       sig.Address().String(), // hex string
					AccountNumber: sig.AccountNumber,
					Pubkey:        consPubKey,
					Sequence:      sig.Sequence,
					Signature:     base64.StdEncoding.EncodeToString(sig.Signature), // encode base64
				}
			}

			sigsBz, err := ex.cdc.MarshalJSON(sigs)
			if err != nil {
				return nil, err
			}

			tempTransaction := &schema.Transaction{
				Height:     tx.Height,
				TxHash:     tx.Hash.String(),
				Code:       tx.TxResult.Code, // 0 is success
				Messages:   string(msgsBz),
				Signatures: string(sigsBz),
				Memo:       stdTx.Memo,
				GasWanted:  tx.TxResult.GasWanted,
				GasUsed:    tx.TxResult.GasUsed,
				Timestamp:  block.Block.Time,
			}

			transactions = append(transactions, tempTransaction)

		}
	}

	return transactions, nil
}
