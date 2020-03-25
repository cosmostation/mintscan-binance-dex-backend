package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// BlockController passes requests to its respective service
func BlockController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		services.GetBlocks(db, w, r)
	}).Methods("GET")
}
