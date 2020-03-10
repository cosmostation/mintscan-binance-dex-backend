package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/services"

	"github.com/gorilla/mux"

	amino "github.com/tendermint/go-amino"
)

// AssetController passes requests to its respective service
func AssetController(cdc *amino.Codec, client client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssets(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/asset-holders", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssetHolders(client, db, w, r)
	}).Methods("GET")
}
