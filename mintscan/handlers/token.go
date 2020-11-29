package handlers

import (
	"net/http"
	"strconv"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
)

// GetTokens returns assets based upon the request params
func GetTokens(c *gin.Context) {
	limit := 100
	offset := 0

	q := c.Request.URL.Query()
	if len(q["limit"]) > 0 {
		limit, _ = strconv.Atoi(q["limit"][0])
	}

	if len(q["offset"]) > 0 {
		offset, _ = strconv.Atoi(q["offset"][0])
	}

	if limit > 1000 {
		errors.ErrOverMaxLimit(c.Writer, http.StatusUnauthorized)
		return
	}

	tks, _ := s.client.GetTokens(limit, offset)

	models.Respond(c.Writer, tks)
	return
}
