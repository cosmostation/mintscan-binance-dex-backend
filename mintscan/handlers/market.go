package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
)

// GetCoinMarketData returns market data from CoinGecko API
func GetCoinMarketData(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["id"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'id' is not present")
		return
	}

	id := r.URL.Query()["id"][0]

	data, err := s.client.GetCoinMarketData(id)
	if err != nil {
		s.l.Printf("failed to fetch coin market data: %s\n", err)
	}

	marketData := &models.Market{
		Name:              data.Name,
		Symbol:            data.Symbol,
		CurrentPrice:      data.MarketData.CurrentPrice.Usd,
		Currency:          "usd",
		MarketCapRank:     data.MarketCapRank,
		MarketCap:         data.MarketData.MarketCap.Usd,
		PercentChange1H:   data.MarketData.PriceChangePercentage1HInCurrency.Usd,
		PercentChange24H:  data.MarketData.PriceChangePercentage24HInCurrency.Usd,
		PercentChange7D:   data.MarketData.PriceChangePercentage7DInCurrency.Usd,
		TotalVolume:       data.MarketData.TotalVolume.Usd,
		TotalSupply:       data.MarketData.TotalSupply,
		CirculatingSupply: data.MarketData.CirculatingSupply,
		LastUpdated:       data.MarketData.LastUpdated,
	}

	models.Respond(rw, marketData)
	return
}

// GetCoinMarketChartData returns market chart data from CoinGecko API
func GetCoinMarketChartData(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["id"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'id' is not present")
		return
	}

	id := r.URL.Query()["id"][0]

	// Current time and its minus 24 hours
	to := time.Now().UTC()
	from := to.AddDate(0, 0, -1)

	marketChartData, err := s.client.GetCoinMarketChartData(id, fmt.Sprintf("%d", from.Unix()), fmt.Sprintf("%d", to.Unix()))
	if err != nil {
		s.l.Printf("failed to fetch coin market chart data: %s\n", err)
	}

	models.Respond(rw, marketChartData)
	return
}
