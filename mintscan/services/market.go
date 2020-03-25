package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// GetCoinMarketData returns market data from CoinGecko API
func GetCoinMarketData(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["id"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'id' is not present")
		return nil
	}

	id := r.URL.Query()["id"][0]

	data, err := c.CoinMarketData(id)
	if err != nil {
		log.Printf("failed to fetch coin market data: %s\n", err)
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
func GetCoinMarketChartData(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["id"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'id' is not present")
		return nil
	}

	id := r.URL.Query()["id"][0]

	// Current time and its minus 24 hours
	to := time.Now().UTC()
	from := to.AddDate(0, 0, -1)

	marketChartData, err := c.CoinMarketChartData(id, fmt.Sprintf("%d", from.Unix()), fmt.Sprintf("%d", to.Unix()))
	if err != nil {
		log.Printf("failed to fetch coin market chart data: %s\n", err)
	}

	utils.Respond(w, marketChartData)
	return nil
}

// GetCoinList returns market chart data from CoinGecko API
func GetCoinList(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	coinList := make([]models.CoinGeckoCoinList, 0)

	for _, coin := range models.CoinList {
		tempCoinList := &models.CoinGeckoCoinList{
			ID:     coin.ID,
			Symbol: coin.Symbol,
			Name:   coin.Name,
		}

		coinList = append(coinList, *tempCoinList)
	}

	utils.Respond(w, coinList)
	return nil
}
