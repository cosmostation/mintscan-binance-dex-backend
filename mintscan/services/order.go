package services

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"

	"github.com/gorilla/mux"
)

// GetOrders returns order information based up on order id
func GetOrders(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "order id is required")
		return nil
	}

	order, err := c.Order(id)
	if err != nil {
		log.Printf("failed to request order information: %s\n", err)
	}

	utils.Respond(w, order)
	return nil
}
