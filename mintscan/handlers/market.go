package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

// GetCoinMarketData returns market data from CoinGecko API
func GetCoinMarketData(c *gin.Context) {
	q := c.Request.URL.Query()

	if len(q["id"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'id' is not present")
		return
	}

	id := q["id"][0]

	data, err := s.client.GetCoinMarketData(id)
	if err != nil {
		s.l.Printf("failed to fetch coin market data: %s\n", err)
	}

	marketData := &models.Market{
		Name:              data.Name,
		Symbol:            data.Symbol,
		CurrentPrice:      decimal.NewFromFloat(data.MarketData.CurrentPrice.Usd).StringFixedBank(2),
		Currency:          "usd",
		MarketCapRank:     data.MarketCapRank,
		MarketCap:         decimal.NewFromFloat(data.MarketData.MarketCap.Usd).StringFixedBank(0),
		PercentChange1H:   decimal.NewFromFloat(data.MarketData.PriceChangePercentage1HInCurrency.Usd).StringFixedBank(2),
		PercentChange24H:  decimal.NewFromFloat(data.MarketData.PriceChangePercentage24HInCurrency.Usd).StringFixedBank(2),
		PercentChange7D:   decimal.NewFromFloat(data.MarketData.PriceChangePercentage7DInCurrency.Usd).StringFixedBank(2),
		TotalVolume:       decimal.NewFromFloat(data.MarketData.TotalVolume.Usd).StringFixedBank(0),
		TotalSupply:       decimal.NewFromFloat(data.MarketData.TotalSupply).StringFixedBank(0),
		CirculatingSupply: decimal.NewFromFloat(data.MarketData.CirculatingSupply).StringFixedBank(0),
		LastUpdated:       data.MarketData.LastUpdated,
	}

	models.Respond(c.Writer, marketData)
	return
}

// GetCoinMarketChartData returns market chart data from CoinGecko API
func GetCoinMarketChartData(c *gin.Context) {
	q := c.Request.URL.Query()

	if len(q["id"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'id' is not present")
		return
	}

	id := q["id"][0]

	// Current time and its minus 24 hours
	to := time.Now().UTC()
	from := to.AddDate(0, 0, -1)

	marketChartData, err := s.client.GetCoinMarketChartData(id, fmt.Sprintf("%d", from.Unix()), fmt.Sprintf("%d", to.Unix()))
	if err != nil {
		s.l.Printf("failed to fetch coin market chart data: %s\n", err)
	}

	models.Respond(c.Writer, marketChartData)
	return
}
