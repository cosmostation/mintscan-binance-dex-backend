package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// TokenController passes requests to its respective service
func TokenController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/tokens", func(w http.ResponseWriter, r *http.Request) {
		services.GetTokens(client, db, w, r)
	}).Methods("GET")
}
