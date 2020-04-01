package models

import "encoding/json"

// Account defines the structure for account information
type Account struct {
	Address       string          `json:"address"`
	PublicKey     json.RawMessage `json:"public_key"`
	AccountNumber int64           `json:"account_number"`
	Sequence      int64           `json:"sequence"`
	Flags         uint64          `json:"flags"`
	Balances      []struct {
		Symbol string `json:"symbol"`
		Free   string `json:"free"`
		Locked string `json:"locked"`
		Frozen string `json:"frozen"`
	} `json:"balances"`
}

// AccountTxs defines the structure for asset transactions
type AccountTxs struct {
	TxNums  int `json:"txNums"`
	TxArray []struct {
		TxHash        string  `json:"txHash"`
		BlockHeight   int64   `json:"blockHeight"`
		TxType        string  `json:"txType"`
		TimeStamp     int64   `json:"timeStamp"`
		FromAddr      string  `json:"fromAddr"`
		ToAddr        string  `json:"toAddr"`
		Value         float64 `json:"value"`
		TxAsset       string  `json:"txAsset"`
		TxQuoteAsset  string  `json:"txQuoteAsset"`
		TxFee         float64 `json:"txFee"`
		TxAge         int64   `json:"txAge"`
		OrderID       string  `json:"orderId"`
		Data          string  `json:"data,omitempty"`
		Code          int64   `json:"code"`
		Log           string  `json:"log"`
		ConfirmBlocks int64   `json:"confirmBlocks"`
		Memo          string  `json:"memo"`
		Source        int64   `json:"source"`
		HasChildren   int64   `json:"hasChildren"`
	} `json:"txArray"`
}

// ResultAccountTxs defines the structure for response data for AssetTxs
type ResultAccountTxs struct {
	TxNums  int              `json:"txNums"`
	TxArray []AccountTxArray `json:"txArray"`
}

// AccountTxArray wraps ResultAccountTxs TxArray
type AccountTxArray struct {
	BlockHeight   int64          `json:"blockHeight"`
	Code          int64          `json:"code"`
	TxHash        string         `json:"txHash"`
	TxType        string         `json:"txType"`
	TxAsset       string         `json:"txAsset"`
	TxQuoteAsset  string         `json:"txQuoteAsset,omitempty"`
	Value         float64        `json:"value"`
	TxFee         float64        `json:"txFee"`
	FromAddr      string         `json:"fromAddr"`
	ToAddr        string         `json:"toAddr,omitempty"`
	TxAge         int64          `json:"txAge"`
	OrderID       string         `json:"orderId,omitempty"`
	Message       *AccountTxData `json:"message,omitempty"`
	Log           string         `json:"log"`
	ConfirmBlocks int64          `json:"confirmBlocks"`
	Memo          string         `json:"memo"`
	Source        int64          `json:"source"`
	HasChildren   int64          `json:"hasChildren,omitempty"`
	Timestamp     int64          `json:"timeStamp"`
}

// AccountTxData defines the structure for ResultAccountTxs Data
type AccountTxData struct {
	OrderData struct {
		Symbol      string `json:"symbol"`
		OrderType   string `json:"orderType"`
		Side        string `json:"side"`
		Price       string `json:"price"`
		Quantity    string `json:"quantity"`
		TimeInForce string `json:"timeInForce"`
		OrderID     string `json:"orderId"`
	} `json:"orderData"`
}
