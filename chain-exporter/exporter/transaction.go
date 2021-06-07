package exporter

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/schema"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/types"
	evm "github.com/InjectiveLabs/sdk-go/chain/evm/types"
)

type Message struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

// getTxs parses transactions in a block and return transactions.
func (ex *Exporter) getTxs(block *tmctypes.ResultBlock, ignoreLog bool) (transactions []*schema.Transaction, err error) {
	txs, err := ex.client.GetTxs(block)
	if err != nil {
		return []*schema.Transaction{}, err
	}

	if len(txs) <= 0 {
		return []*schema.Transaction{}, nil
	}

	for _, txResult := range txs {
		tx, err := ex.client.TxDecoder()(txResult.Tx)
		if err != nil {
			err := errors.Wrap(err, "failed to decode result Tx bytes")
			return nil, err
		}

		txMsgs := tx.GetMsgs()
		msgs := make([]Message, 0, len(txMsgs))
		var msgsBz []byte

		for _, msg := range txMsgs {
			msgBz := ex.client.JSONMarshaler().MustMarshalJSON(msg)
			msgs = append(msgs, Message{
				Type:  fmt.Sprintf("/%s", proto.MessageName(msg)),
				Value: json.RawMessage(msgBz),
			})
		}
		msgsBz, _ = json.Marshal(msgs)

		var txType string
		var sigs []types.Signature
		var evmTxData []byte
		var evmTxFrom string
		var evmTxFromAccAddr string
		var txMemo string

		txWithExtensions, ok := tx.(authante.HasExtensionOptionsTx)
		if ok {
			opts := txWithExtensions.GetExtensionOptions()
			if len(opts) > 0 {
				if typeURL := opts[0].GetTypeUrl(); typeURL == "/injective.evm.v1beta1.ExtensionOptionsEthereumTx" {
					txType = "evm"
				} else if typeURL == "/injective.types.v1beta1.ExtensionOptionsWeb3Tx" {
					txType = "cosmos-web3"
				}
			}
		}
		if len(txType) == 0 {
			txType = "cosmos"
		}

		if txType == "evm" {
			impl := tx.GetMsgs()[0].(*evm.MsgEthereumTx)

			from, err := impl.VerifySig(impl.ChainID())
			if err != nil {
				return nil, err
			} else {
				evmTxFrom = from.Hex()
				evmTxFromAccAddr = sdk.AccAddress(from.Bytes()).String()
			}

			evmTxData, _ = json.Marshal(impl.Data)
		} else if impl, ok := tx.(authsigning.Tx); ok {
			txMemo = impl.GetMemo()
			txSignatures, _ := impl.GetSignaturesV2()
			txPubKeys, _ := impl.GetPubKeys()
			txSigners := impl.GetSigners()
			sigs = make([]types.Signature, len(txSignatures))

			for i, sig := range txSignatures {
				consPubKey := sdk.GetConsAddress(txPubKeys[i]).String()
				if err != nil {
					return []*schema.Transaction{}, err
				}

				sigs[i] = types.Signature{
					Address:   txSigners[i].String(),
					Pubkey:    consPubKey,
					Sequence:  sig.Sequence,
					Signature: base64.StdEncoding.EncodeToString(sig.Data.(*signing.SingleSignatureData).Signature),
				}
			}
		} else {
			err := errors.Errorf("unknown tx type: %T (expected %s)", tx, txType)
			log.WithField("tx_type", fmt.Sprintf("%T", tx)).WithField("expected_type", txType).Warningln("skipping unknown Tx implementation")
			return nil, err
		}

		sigsBz, err := json.Marshal(sigs)
		if err != nil {
			return []*schema.Transaction{}, err
		}

		eventsBz, err := json.Marshal(txResult.TxResult.Events)
		if err != nil {
			return []*schema.Transaction{}, err
		}

		var txLog string
		txInfo := txResult.TxResult.Info

		if !ignoreLog {
			if resultLog := txResult.TxResult.Log; json.Valid([]byte(resultLog)) {
				txLog = resultLog
			} else if len(txInfo) == 0 {
				txInfo = resultLog
			}
		}

		t := &schema.Transaction{
			TxType:           txType,
			Height:           txResult.Height,
			TxHash:           txResult.Hash.String(),
			Code:             txResult.TxResult.Code,
			Messages:         string(msgsBz),
			Signatures:       string(sigsBz),
			Events:           string(eventsBz),
			Memo:             txMemo,
			Log:              txLog,
			Info:             txInfo,
			EVMTxData:        string(evmTxData),
			EVMTxFrom:        evmTxFrom,
			EVMTxFromAccAddr: evmTxFromAccAddr,
			GasWanted:        txResult.TxResult.GasWanted,
			GasUsed:          txResult.TxResult.GasUsed,
			Timestamp:        block.Block.Time,
		}

		transactions = append(transactions, t)
	}

	return transactions, nil
}
