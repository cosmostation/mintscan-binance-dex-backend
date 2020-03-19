package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/services"

	"github.com/gorilla/mux"

	amino "github.com/tendermint/go-amino"
)

// MarketController passes requests to its respective service
func MarketController(cdc *amino.Codec, client client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/market", func(w http.ResponseWriter, r *http.Request) {
		services.GetCoinMarketData(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/market/chart", func(w http.ResponseWriter, r *http.Request) {
		services.GetCoinMarketChartData(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/market/coin/list", func(w http.ResponseWriter, r *http.Request) {
		services.GetCoinList(client, db, w, r)
	}).Methods("GET")
}
