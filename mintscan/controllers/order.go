package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// OrderController passes requests to its respective service
func OrderController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		services.GetOrders(client, db, w, r)
	}).Methods("GET")
}
