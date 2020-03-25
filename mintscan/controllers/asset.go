package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/services"

	"github.com/gorilla/mux"
)

// AssetController passes requests to its respective service
func AssetController(client *client.Client, db *db.Database, r *mux.Router) {
	r.HandleFunc("/asset", func(w http.ResponseWriter, r *http.Request) {
		services.GetAsset(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssets(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/asset-holders", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssetHolders(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/assets-images", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssetsImages(client, db, w, r)
	}).Methods("GET")
	r.HandleFunc("/assets/txs", func(w http.ResponseWriter, r *http.Request) {
		services.GetAssetTxs(client, db, w, r)
	}).Methods("GET")
}
