package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
)

// GetOrders returns order information based up on order id
func GetOrders(c *gin.Context) {
	id := c.Params.ByName("id")

	if id == "" {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "order id is required")
		return
	}

	order, err := s.client.GetOrder(id)
	if err != nil {
		s.l.Printf("failed to request order information: %s\n", err)
		return
	}

	models.Respond(c.Writer, order)
	return
}
