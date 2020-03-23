package services

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

// GetAssetsChartHistory returns
func GetAssetsChartHistory(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	result := make([]models.AssetChartHistory, 0)

	limit := int(24)

	for _, assetName := range models.ChosenAssetNames {
		asset, err := client.Asset(assetName)
		if err != nil {
			log.Printf("failed to get asset detail information: %s\n", err)
		}

		charts, err := db.QueryAssetChartHistory(assetName, limit)
		if err != nil {
			log.Printf("failed to query asset chart history: %s", err)
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

	utils.Respond(w, result)
	return nil
}
