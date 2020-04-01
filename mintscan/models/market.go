package models

import "time"

// Market defines the structure for market data
// This project uses CoinGecko API
type Market struct {
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	CurrentPrice      float64   `json:"current_price"`
	Currency          string    `json:"currency"`
	MarketCapRank     uint8     `json:"market_cap_rank"`
	MarketCap         float64   `json:"market_cap"`
	PercentChange1H   float64   `json:"percent_change_1h"`
	PercentChange24H  float64   `json:"percent_change_24h"`
	PercentChange7D   float64   `json:"percent_change_7d"`
	TotalVolume       float64   `json:"total_volume"`
	TotalSupply       float64   `json:"total_supply"`
	CirculatingSupply float64   `json:"circulating_supply"`
	LastUpdated       time.Time `json:"last_updated"`
}
