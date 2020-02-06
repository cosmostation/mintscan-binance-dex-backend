package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

// GetBlocks returns blocks based upon the request params
func GetBlocks(db *db.Database, w http.ResponseWriter, r *http.Request) error {
	limit := int(100)
	before := int(-1)
	after := int(-1)
	offset := int(0)

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["before"]) > 0 {
		before, _ = strconv.Atoi(r.URL.Query()["before"][0])
	}

	if len(r.URL.Query()["after"]) > 0 {
		after, _ = strconv.Atoi(r.URL.Query()["after"][0])
	}

	if len(r.URL.Query()["offset"]) > 0 {
		offset, _ = strconv.Atoi(r.URL.Query()["offset"][0])
	}

	if limit > 100 {
		errors.ErrOverMaxLimit(w, http.StatusUnauthorized)
		return nil
	}

	var blocks []schema.Block

	var err error

	switch {
	case before > 0:
		blocks, err = db.QueryBlocks(limit, before, after, offset)
		if err != nil {
			fmt.Println(err)
		}
	case after > 0:
		blocks, err = db.QueryBlocks(limit, before, after, offset)
		if err != nil {
			fmt.Println(err)
		}
	case offset >= 0:
		blocks, err = db.QueryBlocks(limit, before, after, offset)
		if err != nil {
			fmt.Println(err)
		}
	}

	utils.Respond(w, blocks)
	return nil
}
