package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// TxController passes requests to its respective service
func TxController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/txs", func(w http.ResponseWriter, r *http.Request) {
		services.GetTxs(db, w, r)
	}).Methods("GET")
	r.HandleFunc("/txs/{hash}", func(w http.ResponseWriter, r *http.Request) {
		services.GetTxByHash(db, w, r)
	}).Methods("GET")
	r.HandleFunc("/txs", func(w http.ResponseWriter, r *http.Request) {
		services.GetTxsByType(client, db, w, r)
	}).Methods("POST")
}
