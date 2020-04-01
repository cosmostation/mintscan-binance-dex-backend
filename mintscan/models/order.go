package models

import "time"

// Order defines the structure for order information
type Order struct {
	OrderID              string    `json:"orderId"`
	Symbol               string    `json:"symbol"`
	Owner                string    `json:"owner"`
	Price                string    `json:"price"`
	Quantity             string    `json:"quantity"`
	CumulateQuantity     string    `json:"cumulateQuantity"`
	Fee                  string    `json:"fee"`
	OrderCreateTime      time.Time `json:"orderCreateTime"`
	TransactionTime      time.Time `json:"transactionTime"`
	Status               string    `json:"status"`
	TimeInForce          int       `json:"timeInForce"`
	Side                 int       `json:"side"`
	Type                 int       `json:"type"`
	TradeID              string    `json:"tradeId"`
	LastExecutedPrice    string    `json:"lastExecutedPrice"`
	LastExecutedQuantity string    `json:"lastExecutedQuantity"`
	TransactionHash      string    `json:"transactionHash"`
}
