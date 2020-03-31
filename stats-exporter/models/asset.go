package models

import (
	"encoding/json"
)

// ChosenAssetNames define 4 asset names
// that are displayed on the card view on Asset page
var ChosenAssetNames = []string{
	"TUSDB-888",
	"USDSB-1AC",
	"BTCB-1DE",
	"IRIS-D88",
}

// Asset defines the structure for asset detail information
type Asset struct {
	CreateTime      interface{} `json:"createTime"`
	UpdateTime      interface{} `json:"updateTime"`
	ID              int         `json:"id"`
	Asset           string      `json:"asset"`
	MappedAsset     string      `json:"mappedAsset"`
	Name            string      `json:"name"`
	AssetImg        string      `json:"assetImg"`
	Supply          float64     `json:"supply"`
	Price           float64     `json:"price"`
	QuoteUnit       string      `json:"quoteUnit"`
	ChangeRange     float64     `json:"changeRange"`
	Owner           string      `json:"owner"`
	Mintable        int         `json:"mintable"`
	Visible         interface{} `json:"visible"`
	Description     string      `json:"description"`
	AssetCreateTime interface{} `json:"assetCreateTime"`
	Transactions    int         `json:"transactions"`
	Holders         int         `json:"holders"`
	OfficialSiteURL string      `json:"officialSiteUrl"`
	ContactEmail    string      `json:"contactEmail"`
	MediaList       []struct {
		MediaName string `json:"mediaName"`
		MediaURL  string `json:"mediaUrl"`
		MediaImg  string `json:"mediaImg"`
	} `json:"mediaList"`
}

// Assets defines the structure for asset information list
type Assets struct {
	TotalNum      int `json:"totalNum"`
	AssetInfoList []struct {
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
