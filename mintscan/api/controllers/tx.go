package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/services"

	"github.com/gorilla/mux"

	amino "github.com/tendermint/go-amino"
)

// TxController passes requests to its respective service
func TxController(cdc *amino.Codec, client client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/txs", func(w http.ResponseWriter, r *http.Request) {
		services.GetTxs(db, w, r)
	}).Methods("GET")
	r.HandleFunc("/txs", func(w http.ResponseWriter, r *http.Request) {
		services.GetTxsByType(client, db, w, r)
	}).Methods("POST")
}
