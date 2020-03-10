package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

// GetAssets returns assets based upon the request params
func GetAssets(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'page' is not present")
		return nil
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'rows' is not present")
		return nil
	}

	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "'rows' cannot be less than 1")
		return nil
	}

	result, err := client.Assets(page, rows)
	if err != nil {
		fmt.Printf("failed to get asset list: %t\n", err)
	}

	utils.Respond(w, result)
	return nil
}

// GetAssetHolders returns asset holders based upon the request params
func GetAssetHolders(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["asset"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'asset' is not present")
		return nil
	}

	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'page' is not present")
		return nil
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'rows' is not present")
		return nil
	}

	asset := r.URL.Query()["asset"][0]
	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "'rows' cannot be less than 1")
		return nil
	}

	result, err := client.AssetHolders(asset, page, rows)
	if err != nil {
		fmt.Printf("failed to get asset holders list: %t\n", err)
	}

	utils.Respond(w, result)
	return nil
}
