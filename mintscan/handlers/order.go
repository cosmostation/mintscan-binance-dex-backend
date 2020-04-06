package handlers

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"

	"github.com/gorilla/mux"
)

// Order is a order handler
type Order struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewOrder creates a new order handler with the given params
func NewOrder(l *log.Logger, client *client.Client, db *db.Database) *Order {
	return &Order{l, client, db}
}

// GetOrders returns order information based up on order id
func (o *Order) GetOrders(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "order id is required")
		return
	}

	order, err := o.client.Order(id)
	if err != nil {
		o.l.Printf("failed to request order information: %s\n", err)
	}

	utils.Respond(rw, order)
	return
}
