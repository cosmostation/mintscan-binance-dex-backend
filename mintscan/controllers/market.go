package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// MarketController passes requests to its respective service
func MarketController(client *client.Client, db *db.Database, r *mux.Router) {
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
