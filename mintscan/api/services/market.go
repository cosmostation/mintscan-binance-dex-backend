package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
)

// GetCoinMarketData returns market data from CoinGecko API
func GetCoinMarketData(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		errors.ErrRequireIDParam(w, http.StatusBadRequest)
		return nil
	}

	data, err := client.CoinMarketData(int(0))
	if err != nil {
		fmt.Printf("failed to fetch coin market data: %t\n", err)
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

	utils.Respond(w, marketData)
	return nil
}

// GetCoinMarketChartData returns market chart data from CoinGecko API
func GetCoinMarketChartData(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		errors.ErrRequireIDParam(w, http.StatusBadRequest)
		return nil
	}

	// Current time and its minus 24 hours
	to := time.Now().UTC()
	from := to.AddDate(0, 0, -1)

	// Convert from unix timestamp to string
	toStr := fmt.Sprintf("%d", to.Unix())
	fromStr := fmt.Sprintf("%d", from.Unix())

	fmt.Println(toStr)
	fmt.Println(fromStr)

	marketChartData, err := client.CoinMarketChartData(int(0), int(0), int(0))
	// marketChartData, err := client.CoinMarketChartData(id, fromStr, toStr)
	if err != nil {
		fmt.Printf("failed to fetch coin market chart data: %t\n", err)
	}

	utils.Respond(w, marketChartData)
	return nil
}
