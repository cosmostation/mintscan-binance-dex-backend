package schema

import "time"

// TransactionInfo represents the information a transaction contains
type TransactionInfo struct {
	ID         int32     `json:"id" sql:",pk"`
	Height     int32     `json:"height" sql:",notnull"`
	TxHash     string    `json:"tx_hash" sql:",notnull"`
	GasWanted  int32     `json:"gas_wanted"`
	GasUsed    int32     `json:"gas_used"`
	Events     struct{}  `json:"events"`
	Messages   struct{}  `json:"messages" sql:",notnull"`
	Fee        struct{}  `json:"fee" sql:",notnull"`
	Signatures struct{}  `json:"signautures" sql:",notnull"`
	Memo       string    `json:"memo"`
	Timestamp  time.Time `json:"timestamp" sql:"default:now()"`
}
