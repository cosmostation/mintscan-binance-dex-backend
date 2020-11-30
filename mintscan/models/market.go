package models

import "time"

// Market defines the structure for market data
// This project uses CoinGecko API
type Market struct {
	Name              string    `json:"name"`
	Symbol            string    `json:"symbol"`
	CurrentPrice      string    `json:"current_price"`
	Currency          string    `json:"currency"`
	MarketCapRank     int       `json:"market_cap_rank"`
	MarketCap         string    `json:"market_cap"`
	PercentChange1H   string    `json:"percent_change_1h"`
	PercentChange24H  string    `json:"percent_change_24h"`
	PercentChange7D   string    `json:"percent_change_7d"`
	TotalVolume       string    `json:"total_volume"`
	TotalSupply       string    `json:"total_supply"`
	CirculatingSupply string    `json:"circulating_supply"`
	LastUpdated       time.Time `json:"last_updated"`
}
