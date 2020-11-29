package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
)

// GetAsset returns asset based upon the request params
func GetAsset(c *gin.Context) {
	q := c.Request.URL.Query()
	if len(q["asset"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'asset' is not present")
		return
	}

	asset := q["asset"][0]

	result, err := s.client.GetAsset(asset)
	if err != nil {
		s.l.Printf("failed to get asset detail information: %s\n", err)
	}

	models.Respond(c.Writer, result)
	return
}

// GetAssets returns assets based upon the request params
func GetAssets(c *gin.Context) {
	onlyPrice := "false" // default is false, when true it only show assets price information

	q := c.Request.URL.Query()
	if len(q["page"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(q["rows"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'rows' is not present")
		return
	}

	if len(q["only_price"]) > 0 {
		onlyPrice = q["only_price"][0]
	}

	page, _ := strconv.Atoi(q["page"][0])
	rows, _ := strconv.Atoi(q["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 1000 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be greater than 100")
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

		models.Respond(c.Writer, result)
		return
	}

	models.Respond(c.Writer, assets)
	return
}

// GetAssetHolders returns asset holders based upon the request params
func GetAssetHolders(c *gin.Context) {
	q := c.Request.URL.Query()

	if len(q["asset"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'asset' is not present")
		return
	}

	if len(q["page"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(q["rows"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'rows' is not present")
		return
	}

	asset := q["asset"][0]
	page, _ := strconv.Atoi(q["page"][0])
	rows, _ := strconv.Atoi(q["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 100 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be greater than 100")
		return
	}

	result, err := s.client.GetAssetHolders(asset, page, rows)
	if err != nil {
		s.l.Printf("failed to get asset holders list: %s\n", err)
	}

	models.Respond(c.Writer, result)
	return
}

// GetAssetsImages returns images of all assets
func GetAssetsImages(c *gin.Context) {
	q := c.Request.URL.Query()

	if len(q["page"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(q["rows"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'rows' is not present")
		return
	}

	page, _ := strconv.Atoi(q["page"][0])
	rows, _ := strconv.Atoi(q["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 100 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be greater than 100")
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

	models.Respond(c.Writer, result)
	return
}

// GetAssetTxs returns asset txs
func GetAssetTxs(c *gin.Context) {
	q := c.Request.URL.Query()

	if len(q["txAsset"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'txAsset' is not present")
		return
	}

	if len(q["page"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(q["rows"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'rows' is not present")
		return
	}

	txAsset := q["txAsset"][0]
	page, _ := strconv.Atoi(q["page"][0])
	rows, _ := strconv.Atoi(q["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 100 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be greater than 100")
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

	models.Respond(c.Writer, result)
	return
}

// GetAssetsMiniTokens returns a list of mini tokens based upon the request params.
func GetAssetsMiniTokens(c *gin.Context) {
	q := c.Request.URL.Query()
	onlyPrice := "false" // default is false, when true it only show assets price information

	if len(q["page"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'page' is not present")
		return
	}

	if len(q["rows"]) <= 0 {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "'rows' is not present")
		return
	}

	if len(q["only_price"]) > 0 {
		onlyPrice = q["only_price"][0]
	}

	page, _ := strconv.Atoi(q["page"][0])
	rows, _ := strconv.Atoi(q["rows"][0])

	if rows < 1 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 1000 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be greater than 100")
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

		models.Respond(c.Writer, result)
		return
	}

	models.Respond(c.Writer, assets)
	return
}
