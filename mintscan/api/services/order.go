package services

import (
	"fmt"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"

	"github.com/gorilla/mux"
)

// GetOrders returns order information based up on order id
func GetOrders(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "order id is required")
		return nil
	}

	order, err := client.Order(id)
	if err != nil {
		fmt.Printf("failed to request order information: %s\n", err)
	}

	utils.Respond(w, order)
	return nil
}
