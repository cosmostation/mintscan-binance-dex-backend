package schema

import "time"

// TransactionInfo represents the information in a transaction
type TransactionInfo struct {
	ID         int32     `json:"id" sql:",pk"`
	Height     int32     `json:"height" sql:",notnull"`
	TxHash     string    `json:"tx_hash" sql:",notnull,unique"`
	Events     struct{}  `json:"events"`
	Messages   struct{}  `json:"messages" sql:",notnull"`
	Fee        struct{}  `json:"fee" sql:",notnull"`
	Signatures struct{}  `json:"signautures" sql:",notnull"`
	Memo       string    `json:"memo"`
	GasWanted  int32     `json:"gas_wanted"`
	GasUsed    int32     `json:"gas_used"`
	Timestamp  time.Time `json:"timestamp" sql:"default:now()"`
}
