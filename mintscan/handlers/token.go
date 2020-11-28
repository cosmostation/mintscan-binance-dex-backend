package handlers

import (
	"net/http"
	"strconv"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
)

// GetTokens returns assets based upon the request params
func GetTokens(rw http.ResponseWriter, r *http.Request) {
	limit := 100
	offset := 0

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["offset"]) > 0 {
		offset, _ = strconv.Atoi(r.URL.Query()["offset"][0])
	}

	if limit > 1000 {
		errors.ErrOverMaxLimit(rw, http.StatusUnauthorized)
		return
	}

	tks, _ := s.client.GetTokens(limit, offset)

	models.Respond(rw, tks)
	return
}
