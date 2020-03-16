package exporter

import "github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

// getAssetInfoList1H fetches asset information list and return them
func (ex *Exporter) getAssetInfoList1H() ([]*schema.StatAssetInfoList1H, error) {
	assetInfoList := make([]*schema.StatAssetInfoList1H, 0)

	page := int(1)
	rows := int(200)

	assets, err := ex.client.AssetInfoList(page, rows)
	if err != nil {
		return assetInfoList, err
	}

	for _, asset := range assets.AssetInfoList {
		tempAssetInfoList := &schema.StatAssetInfoList1H{
			TotalNum:        assets.TotalNum,
			Name:            asset.Name,
			Asset:           asset.Asset,
			Owner:           asset.Owner,
			Price:           asset.Price,
			Currency:        asset.QuoteUnit,
			ChangeRange:     asset.ChangeRange,
			Supply:          asset.Supply,
			Marketcap:       asset.Supply * asset.Price,
			AssetImg:        asset.AssetImg,
			AssetCreateTime: asset.AssetCreateTime,
		}

		assetInfoList = append(assetInfoList, tempAssetInfoList)
	}

	return assetInfoList, nil
}

// getAssetInfoList24H fetches asset information list and return them
func (ex *Exporter) getAssetInfoList24H() ([]*schema.StatAssetInfoList24H, error) {
	assetInfoList := make([]*schema.StatAssetInfoList24H, 0)

	page := int(1)
	rows := int(200)

	assets, err := ex.client.AssetInfoList(page, rows)
	if err != nil {
		return assetInfoList, err
	}

	for _, asset := range assets.AssetInfoList {
		tempAssetInfoList := &schema.StatAssetInfoList24H{
			TotalNum:        assets.TotalNum,
			Name:            asset.Name,
			Asset:           asset.Asset,
			Owner:           asset.Owner,
			Price:           asset.Price,
			Currency:        asset.QuoteUnit,
			ChangeRange:     asset.ChangeRange,
			Supply:          asset.Supply,
			Marketcap:       asset.Supply * asset.Price,
			AssetImg:        asset.AssetImg,
			AssetCreateTime: asset.AssetCreateTime,
		}

		assetInfoList = append(assetInfoList, tempAssetInfoList)
	}

	return assetInfoList, nil
}
