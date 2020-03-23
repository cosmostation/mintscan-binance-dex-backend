package services

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

/*
	1. 정해진 4개의 차트를 던져준다
	[
		{
			name: "",
			asset: "",
			.....
			prices: [

			]
		},
		{
			name: "",
			asset: "",
			.....
			prices: [

			]
		}
	]
*/

// GetAssetsChartHistory returns
func GetAssetsChartHistory(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {

	asset := "USDSB-1AC"

	result, err := db.QueryAssetChartHistory(asset)
	if err != nil {
		log.Printf("failed to query asset chart history: %s", err)
	}

	utils.Respond(w, result)
	return nil
}
