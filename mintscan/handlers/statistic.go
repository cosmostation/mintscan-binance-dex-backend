package handlers

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// Statistic is a statistic handler
type Statistic struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewStatistic creates a new statistic handler with the given params
func NewStatistic(l *log.Logger, client *client.Client, db *db.Database) *Statistic {
	return &Statistic{l, client, db}
}

// GetAssetsChartHistory returns
func (s *Statistic) GetAssetsChartHistory(rw http.ResponseWriter, r *http.Request) {
	result := make([]models.AssetChartHistory, 0)

	limit := int(24)

	for _, assetName := range models.ChosenAssetNames {
		asset, err := s.client.Asset(assetName)
		if err != nil {
			s.l.Printf("failed to get asset detail information: %s\n", err)
		}

		charts, err := s.db.QueryAssetChartHistory(assetName, limit)
		if err != nil {
			s.l.Printf("failed to query asset chart history: %s", err)
		}

		prices := make([]models.Prices, 0)

		for _, chart := range charts {
			tempPrices := &models.Prices{
				Price:     chart.Price,
				Timestamp: chart.Timestamp,
			}

			prices = append(prices, *tempPrices)
		}

		tempResult := &models.AssetChartHistory{
			Name:         asset.Name,
			Asset:        asset.Asset,
			MappedAsset:  asset.MappedAsset,
			CurrentPrice: asset.Price,
			QuoteUnit:    asset.QuoteUnit,
			ChangeRange:  asset.ChangeRange,
			Supply:       asset.Supply,
			Marketcap:    asset.Price * asset.Supply,
			AssetImage:   asset.AssetImg,
			Prices:       prices,
		}

		result = append(result, *tempResult)
	}

	utils.Respond(rw, result)
	return
}
