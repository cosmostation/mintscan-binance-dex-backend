package controllers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/services"

	"github.com/gorilla/mux"

	"github.com/binance-chain/go-sdk/client/rpc"

	amino "github.com/tendermint/go-amino"
)

// BlockController passes requests to its respective service
func BlockController(cdc *amino.Codec, cfg *config.Config, db *db.Database, r *mux.Router, rpcClient rpc.Client) {
	r.HandleFunc("/blocks", func(w http.ResponseWriter, r *http.Request) {
		services.GetBlocks(db, w, r)
	}).Methods("GET")
}
