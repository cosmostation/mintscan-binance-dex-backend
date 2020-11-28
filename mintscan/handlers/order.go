package handlers

import (
	"net/http"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"

	"github.com/gorilla/mux"
)

// GetOrders returns order information based up on order id
func GetOrders(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "order id is required")
		return
	}

	order, err := s.client.GetOrder(id)
	if err != nil {
		s.l.Printf("failed to request order information: %s\n", err)
		return
	}

	models.Respond(rw, order)
	return
}
