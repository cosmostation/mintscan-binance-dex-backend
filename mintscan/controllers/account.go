package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"
	"github.com/gorilla/mux"
)

// AccountController passes requests to its respective service
func AccountController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/account/{address}", func(w http.ResponseWriter, r *http.Request) {
		services.GetAccount(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/account/txs/{address}", func(w http.ResponseWriter, r *http.Request) {
		services.GetAccountTxs(client, db, w, r)
	}).Methods("GET")
}
