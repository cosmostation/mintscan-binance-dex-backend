package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// StatsController passes requests to its respective service
func StatsController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/stats/assets/chart", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssetsChartHistory(client, db, w, r)
	}).Methods("GET")
}
