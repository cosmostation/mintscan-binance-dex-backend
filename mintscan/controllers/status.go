package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// StatusController passes requests to its respective service
func StatusController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		services.GetStatus(client, db, w, r)
	}).Methods("GET")
}
