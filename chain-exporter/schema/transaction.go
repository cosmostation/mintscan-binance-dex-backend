package schema

import "time"

// Transaction defines the structure for transaction information.
type Transaction struct {
	ID               int32     `json:"id" sql:",pk"`
	TxType           string    `json:"tx_type" sql:",notnull"`
	Height           int64     `json:"height" sql:",notnull"`
	TxHash           string    `json:"tx_hash" sql:",notnull,unique"`
	Code             uint32    `json:"code"  sql:",notnull"`
	Messages         string    `json:"messages" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Signatures       string    `json:"signatures" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Log              string    `json:"log" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Info             string    `json:"info"`
	Memo             string    `json:"memo"`
	Events           string    `json:"events" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	EVMTxData        string    `json:"evm_tx_data" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	EVMTxFrom        string    `json:"evm_tx_from"`
	EVMTxFromAccAddr string    `json:"evm_tx_from_acc"`
	GasWanted        int64     `json:"gas_wanted" sql:"default:0"`
	GasUsed          int64     `json:"gas_used" sql:"default:0"`
	Timestamp        time.Time `json:"timestamp" sql:"default:now()"`
}

// NewTransaction returns a new Transaction.
func NewTransaction(t Transaction) *Transaction {
	return &Transaction{
		Height:           t.Height,
		TxType:           t.TxType,
		TxHash:           t.TxHash,
		Code:             t.Code,
		Messages:         t.Messages,
		Signatures:       t.Signatures,
		Log:              t.Log,
		Info:             t.Info,
		Memo:             t.Memo,
		Events:           t.Events,
		EVMTxData:        t.EVMTxData,
		EVMTxFrom:        t.EVMTxFrom,
		EVMTxFromAccAddr: t.EVMTxFromAccAddr,
		GasWanted:        t.GasWanted,
		GasUsed:          t.GasUsed,
		Timestamp:        t.Timestamp,
	}
}
