package models

import "time"

// Market represents market data
type Market struct {
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	CurrentPrice      float64   `json:"current_price"`
	Currency          string    `json:"currency"`
	MarketCapRank     uint8     `json:"market_cap_rank"`
	PercentChange1H   float64   `json:"percent_change_1h"`
	PercentChange24H  float64   `json:"percent_change_24h"`
	PercentChange7D   float64   `json:"percent_change_7d"`
	PercentChange30D  float64   `json:"percent_change_30d"`
	TotalVolume       uint64    `json:"total_volume"`
	CirculatingSupply float64   `json:"circulating_supply"`
	LastUpdated       time.Time `json:"last_updated"`
	PriceStats        []Price   `json:"price_stats"`
}

// Price represents BNB price
type Price struct {
	Price float64   `json:"price"`
	Time  time.Time `json:"time"`
}
