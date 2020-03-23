package services

import (
	"encoding/json"
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
		log.Printf("failed to get asset detail information: %s\n", err)
	}

	utils.Respond(w, result)
	return nil
}

// GetAssets returns assets based upon the request params
func GetAssets(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	onlyPrice := "false" // default is false, when true it only show assets price information

	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'page' is not present")
		return nil
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "'rows' is not present")
		return nil
	}

	if len(r.URL.Query()["only_price"]) > 0 {
		onlyPrice = r.URL.Query()["only_price"][0]
	}

	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "'rows' cannot be less than 1")
		return nil
	}

	assets, err := client.Assets(page, rows)
	if err != nil {
		log.Printf("failed to get asset list: %s\n", err)
	}

	if onlyPrice == "true" {
		assetInfoList := make([]models.ResultAssetInfoList, 0)

		for _, asset := range assets.AssetInfoList {
			tempAssetInfoList := &models.ResultAssetInfoList{
				Asset:       asset.Asset,
				MappedAsset: asset.MappedAsset,
				Name:        asset.Name,
				Price:       asset.Price,
				QuoteUnit:   asset.QuoteUnit,
			}

			assetInfoList = append(assetInfoList, *tempAssetInfoList)
		}

		result := &models.ResultAssetInfo{
			TotalNum:      assets.TotalNum,
			AssetInfoList: assetInfoList,
		}

		utils.Respond(w, result)
		return nil
	}

	utils.Respond(w, assets)
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
		log.Printf("failed to get asset holders list: %s\n", err)
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
		log.Printf("failed to get asset list: %s\n", err)
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
		log.Printf("failed to get asset list: %s\n", err)
	}

	txArray := make([]models.TxArray, 0)

	for _, tx := range assetTxs.TxArray {
		var toAddr string
		if tx.ToAddr != "" {
			toAddr = tx.ToAddr
		}

		tempTxArray := &models.TxArray{
			BlockHeight:   tx.BlockHeight,
			TxHash:        tx.TxHash,
			Code:          tx.Code,
			TxType:        tx.TxType,
			TxAsset:       tx.TxAsset,
			TxQuoteAsset:  tx.TxQuoteAsset,
			Value:         tx.Value,
			TxFee:         tx.TxFee,
			TxAge:         tx.TxAge,
			FromAddr:      tx.FromAddr,
			ToAddr:        toAddr,
			Log:           tx.Log,
			ConfirmBlocks: tx.ConfirmBlocks,
			Memo:          tx.Memo,
			Source:        tx.Source,
			Timestamp:     tx.TimeStamp,
		}

		// txType TRANSFER shouldn't throw message data
		var data models.AssetTxData
		if tx.Data != "" {
			err = json.Unmarshal([]byte(tx.Data), &data)
			if err != nil {
				log.Printf("failed to unmarshal AssetTxData: %s", err)
			}

			tempTxArray.Message = &data
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
