package types

import (
	"encoding/json"
)

// Asset defines the structure for asset information list
type Asset struct {
	TotalNum  int `json:"totalNum"`
	AssetList []struct {
		CreateTime      json.RawMessage `json:"createTime"`
		UpdateTime      json.RawMessage `json:"updateTime"`
		ID              int32           `json:"id"`
		Asset           string          `json:"asset"`
		MappedAsset     string          `json:"mappedAsset"`
		Name            string          `json:"name"`
		AssetImg        string          `json:"assetImg"`
		Supply          float64         `json:"supply"`
		Price           float64         `json:"price"`
		QuoteUnit       string          `json:"quoteUnit"`
		ChangeRange     float64         `json:"changeRange"`
		Owner           string          `json:"owner"`
		Mintable        int32           `json:"mintable"`
		Visible         json.RawMessage `json:"visible"`
		Description     json.RawMessage `json:"description"`
		AssetCreateTime int64           `json:"assetCreateTime"`
	} `json:"assetInfoList"`
}
