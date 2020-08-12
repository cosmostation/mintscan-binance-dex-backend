package schema

import "time"

// Transaction defines the structure for transaction information.
type Transaction struct {
	ID         int32     `json:"id" sql:",pk"`
	Height     int64     `json:"height" sql:",notnull"`
	TxHash     string    `json:"tx_hash" sql:",notnull,unique"`
	Code       uint32    `json:"code"  sql:",notnull"` // https://docs.binance.org/exchange-integration.html#important-ensuring-transaction-finality
	Messages   string    `json:"messages" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Signatures string    `json:"signautures" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Memo       string    `json:"memo"`
	GasWanted  int64     `json:"gas_wanted" sql:"default:0"`
	GasUsed    int64     `json:"gas_used" sql:"default:0"`
	Timestamp  time.Time `json:"timestamp" sql:"default:now()"`
}

// NewTransaction returns a new Transaction.
func NewTransaction(t Transaction) *Transaction {
	return &Transaction{
		Height:     t.Height,
		TxHash:     t.TxHash,
		Code:       t.Code,
		Messages:   t.Messages,
		Signatures: t.Signatures,
		Memo:       t.Memo,
		GasWanted:  t.GasWanted,
		GasUsed:    t.GasUsed,
		Timestamp:  t.Timestamp,
	}
}
