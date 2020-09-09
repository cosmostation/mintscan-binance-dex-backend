package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
)

// GetAsset returns asset based upon the request params
func GetAsset(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["asset"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'asset' is not present")
		return
	}

	asset := r.URL.Query()["asset"][0]

	result, err := s.client.GetAsset(asset)
	if err != nil {
		s.l.Printf("failed to get asset detail information: %s\n", err)
	}

	models.Respond(rw, result)
	return
}

// GetAssets returns assets based upon the request params
func GetAssets(rw http.ResponseWriter, r *http.Request) {
	onlyPrice := "false" // default is false, when true it only show assets price information

	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'rows' is not present")
		return
	}

	if len(r.URL.Query()["only_price"]) > 0 {
		onlyPrice = r.URL.Query()["only_price"][0]
	}

	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 1000 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 100")
		return
	}

	assets, err := s.client.GetAssets(page, rows)
	if err != nil {
		s.l.Printf("failed to get asset list: %s\n", err)
	}

	if onlyPrice == "true" {
		assetInfoList := make([]models.ResultAssetInfoList, 0)

		for _, asset := range assets.AssetInfoList {
			assetInfo := &models.ResultAssetInfoList{
				Asset:       asset.Asset,
				MappedAsset: asset.MappedAsset,
				Name:        asset.Name,
				Price:       asset.Price,
				QuoteUnit:   asset.QuoteUnit,
			}

			assetInfoList = append(assetInfoList, *assetInfo)
		}

		result := &models.ResultAssetInfo{
			TotalNum:      assets.TotalNum,
			AssetInfoList: assetInfoList,
		}

		models.Respond(rw, result)
		return
	}

	models.Respond(rw, assets)
	return
}

// GetAssetHolders returns asset holders based upon the request params
func GetAssetHolders(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["asset"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'asset' is not present")
		return
	}

	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'rows' is not present")
		return
	}

	asset := r.URL.Query()["asset"][0]
	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 100 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 100")
		return
	}

	result, err := s.client.GetAssetHolders(asset, page, rows)
	if err != nil {
		s.l.Printf("failed to get asset holders list: %s\n", err)
	}

	models.Respond(rw, result)
	return
}

// GetAssetsImages returns images of all assets
func GetAssetsImages(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'rows' is not present")
		return
	}

	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 100 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 100")
		return
	}

	assets, err := s.client.GetAssets(page, rows)
	if err != nil {
		s.l.Printf("failed to get asset list: %s\n", err)
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

	models.Respond(rw, result)
	return
}

// GetAssetTxs returns asset txs
func GetAssetTxs(rw http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()["txAsset"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'txAsset' is not present")
		return
	}

	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'rows' is not present")
		return
	}

	txAsset := r.URL.Query()["txAsset"][0]
	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 100 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 100")
		return
	}

	assetTxs, err := s.client.GetAssetTxs(txAsset, page, rows)
	if err != nil {
		s.l.Printf("failed to get asset list: %s\n", err)
	}

	txArray := make([]models.AssetTxArray, 0)

	for _, tx := range assetTxs.TxArray {
		var toAddr string
		if tx.ToAddr != "" {
			toAddr = tx.ToAddr
		}

		tempTxArray := &models.AssetTxArray{
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
				s.l.Printf("failed to unmarshal AssetTxData: %s", err)
			}

			tempTxArray.Message = &data
		}

		txArray = append(txArray, *tempTxArray)
	}

	result := &models.ResultAssetTxs{
		TxNums:  assetTxs.TxNums,
		TxArray: txArray,
	}

	models.Respond(rw, result)
	return
}

// GetAssetsMiniTokens returns a list of mini tokens based upon the request params.
func GetAssetsMiniTokens(rw http.ResponseWriter, r *http.Request) {
	onlyPrice := "false" // default is false, when true it only show assets price information

	if len(r.URL.Query()["page"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(r.URL.Query()["rows"]) <= 0 {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "'rows' is not present")
		return
	}

	if len(r.URL.Query()["only_price"]) > 0 {
		onlyPrice = r.URL.Query()["only_price"][0]
	}

	page, _ := strconv.Atoi(r.URL.Query()["page"][0])
	rows, _ := strconv.Atoi(r.URL.Query()["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 1000 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 100")
		return
	}

	assets, err := s.client.GetMiniTokens(page, rows)
	if err != nil {
		s.l.Printf("failed to get mini tokens list: %s\n", err)
	}

	if onlyPrice == "true" {
		assetInfoList := make([]models.ResultAssetInfoList, 0)

		for _, asset := range assets.AssetInfoList {
			assetInfo := &models.ResultAssetInfoList{
				Asset:       asset.Asset,
				MappedAsset: asset.MappedAsset,
				Name:        asset.Name,
				Price:       asset.Price,
				QuoteUnit:   asset.QuoteUnit,
			}

			assetInfoList = append(assetInfoList, *assetInfo)
		}

		result := &models.ResultAssetInfo{
			TotalNum:      assets.TotalNum,
			AssetInfoList: assetInfoList,
		}

		models.Respond(rw, result)
		return
	}

	models.Respond(rw, assets)
	return
}
