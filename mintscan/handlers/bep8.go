package handlers

import (
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
)

// GetMiniTokens returns a list of mini tokens based upon the request params.
func GetMiniTokens(rw http.ResponseWriter, r *http.Request) {
	limit := 100

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if limit > 200 {
		errors.ErrOverMaxLimit(rw, http.StatusUnauthorized)
		return
	}

	assets, err := s.client.GetMiniTokens(limit)
	if err != nil {
		s.l.Printf("failed to get asset list: %s\n", err)
		return
	}

	models.Respond(rw, assets)
	return
}
