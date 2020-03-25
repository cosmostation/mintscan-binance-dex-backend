package services

import (
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// GetTokens returns assets based upon the request params
func GetTokens(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	limit := 100
	offset := 0

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["offset"]) > 0 {
		offset, _ = strconv.Atoi(r.URL.Query()["offset"][0])
	}

	if limit > 1000 {
		errors.ErrOverMaxLimit(w, http.StatusUnauthorized)
		return nil
	}

	tks, _ := c.Tokens(limit, offset)

	utils.Respond(w, tks)
	return nil
}
