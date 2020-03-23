package services

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
	"github.com/gorilla/mux"
)

// GetAccount returns account information
func GetAccount(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "address is required")
		return nil
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "address is invalid")
		return nil
	}

	account, err := client.Account(address)
	if err != nil {
		log.Printf("failed to request account information: %s\n", err)
	}

	utils.Respond(w, account)
	return nil
}
