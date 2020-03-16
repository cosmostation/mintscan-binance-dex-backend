package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

// GetAsset returns asset based upon the request params
func GetAsset(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["asset"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'asset' is not present")
		return nil
	}

	asset := r.URL.Query()["asset"][0]

	result, err := client.Asset(asset)
	if err != nil {
		log.Printf("failed to get asset detail information: %t\n", err)
	}

	utils.Respond(w, result)
	return nil
}

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
		log.Printf("failed to get asset list: %t\n", err)
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
		log.Printf("failed to get asset holders list: %t\n", err)
	}

	utils.Respond(w, result)
	return nil
}

// GetAssetsImages returns images of all assets
func GetAssetsImages(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
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

	assets, err := client.Assets(page, rows)
	if err != nil {
		log.Printf("failed to get asset list: %t\n", err)
	}

	imageList := make([]models.ImageList, 0)

	for _, asset := range assets.AssetInfoList {
		tempList := &models.ImageList{
			Asset:      asset.Asset,
			Name:       asset.Name,
			AssetImage: asset.AssetImg,
		}

		imageList = append(imageList, *tempList)
	}

	result := &models.ResultAssetsImages{
		TotalNum:  assets.TotalNum,
		ImageList: imageList,
	}

	utils.Respond(w, result)
	return nil
}

// GetAssetTxs returns asset txs
func GetAssetTxs(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	if len(r.URL.Query()["txAsset"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'txAsset' is not present")
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

	txAsset := r.URL.Query()["txAsset"][0]
	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "'rows' cannot be less than 1")
		return nil
	}

	assetTxs, err := client.AssetTxs(txAsset, page, rows)
	if err != nil {
		log.Printf("failed to get asset list: %t\n", err)
	}

	txArray := make([]models.TxArray, 0)

	for _, tx := range assetTxs.TxArray {
		var data models.AssetTxData
		err = json.Unmarshal([]byte(tx.Data), &data)
		if err != nil {
			fmt.Printf("failed to unmarshal AssetTxData: %s", err)
		}

		tempTxArray := &models.TxArray{
			TxHash:        tx.TxHash,
			BlockHeight:   tx.BlockHeight,
			Code:          tx.Code,
			TxType:        tx.TxType,
			TxFee:         tx.TxFee,
			Timestamp:     tx.TimeStamp,
			FromAddr:      tx.FromAddr,
			Value:         tx.Value,
			TxAsset:       tx.TxAsset,
			TxQuoteAsset:  tx.TxQuoteAsset,
			TxAge:         tx.TxAge,
			OrderID:       tx.OrderID,
			Message:       data,
			Log:           tx.Log,
			ConfirmBlocks: tx.ConfirmBlocks,
			Memo:          tx.Memo,
			Source:        tx.Source,
			HasChildren:   tx.HasChildren,
		}

		txArray = append(txArray, *tempTxArray)
	}

	result := &models.ResultAssetTxs{
		TxNums:  assetTxs.TxNums,
		TxArray: txArray,
	}

	utils.Respond(w, result)
	return nil
}

// GetAssetChart returns asset chart
func GetAssetChart(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	log.Println("GetAssetChart")

	// ResultAssetChart
	// Price, Currency, Marketcap, ChangeRange,
	// Prices
	return nil
}
