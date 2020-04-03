package handlers

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

// Market is a market handler
type Market struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewMarket creates a new market handler with the given params
func NewMarket(l *log.Logger, client *client.Client, db *db.Database) *Market {
	return &Market{l, client, db}
}

// GetCoinMarketData returns market data from CoinGecko API
func (m *Market) GetCoinMarketData(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["id"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'id' is not present")
		return
	}

	id := r.URL.Query()["id"][0]

	data, err := m.client.CoinMarketData(id)
	if err != nil {
		m.l.Printf("failed to fetch coin market data: %s\n", err)
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

	utils.Respond(rw, marketData)
	return
}

// GetCoinMarketChartData returns market chart data from CoinGecko API
func (m *Market) GetCoinMarketChartData(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["id"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'id' is not present")
		return
	}

	id := r.URL.Query()["id"][0]

	// Current time and its minus 24 hours
	to := time.Now().UTC()
	from := to.AddDate(0, 0, -1)

	marketChartData, err := m.client.CoinMarketChartData(id, fmt.Sprintf("%d", from.Unix()), fmt.Sprintf("%d", to.Unix()))
	if err != nil {
		m.l.Printf("failed to fetch coin market chart data: %s\n", err)
	}

	utils.Respond(rw, marketChartData)
	return
}
