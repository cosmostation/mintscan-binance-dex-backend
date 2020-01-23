package schema

import "time"

// TransactionInfo represents the information in a transaction
type TransactionInfo struct {
	ID         int32     `json:"id" sql:",pk"`
	Height     int64     `json:"height" sql:",notnull"`
	TxHash     string    `json:"tx_hash" sql:",notnull,unique"`
	GasWanted  int64     `json:"gas_wanted"`
	GasUsed    int64     `json:"gas_used"`
	Messages   string    `json:"messages" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Fee        string    `json:"fee" sql:"type:jsonb, notnull default: '{}'::jsonb"`
	Signatures string    `json:"signautures" sql:"type:jsonb, notnull, default: '[]'::jsonb"`
	Memo       string    `json:"memo"`
	Logs       string    `json:"logs" sql:"type:jsonb, default: '[]'::jsonb"`
	Events     string    `json:"events" sql:"type:jsonb, default: '[]'::jsonb"`
	Timestamp  time.Time `json:"timestamp" sql:"default:now()"`
}
