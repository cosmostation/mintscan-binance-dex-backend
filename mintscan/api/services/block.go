package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
)

// GetBlocks returns latest blocks
func GetBlocks(db *db.Database, w http.ResponseWriter, r *http.Request) error {
	limit := int(100)
	afterBlock := int(1)

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	// check max limit
	if limit > 100 {
		errors.ErrOverMaxLimit(w, http.StatusUnauthorized)
		return nil
	}

	if len(r.URL.Query()["afterBlock"]) > 0 {
		afterBlock, _ = strconv.Atoi(r.URL.Query()["afterBlock"][0])
	}

	fmt.Println("limit: ", limit)
	fmt.Println("afterBlock: ", afterBlock)

	// utils.Respond(w, result)
	return nil
}
