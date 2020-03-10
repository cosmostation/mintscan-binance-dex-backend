package models

// Assets represents asset detail information
type Assets struct {
	TotalNum      int `json:"totalNum"`
	AssetInfoList []struct {
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
		Description     interface{} `json:"description"`
		AssetCreateTime int64       `json:"assetCreateTime"`
	} `json:"assetInfoList"`
}

// AssetHolders represents asset holders list
type AssetHolders struct {
	TotalNum       int `json:"totalNum"`
	AddressHolders []struct {
		Address    string      `json:"address"`
		Quantity   float64     `json:"quantity"`
		Percentage float64     `json:"percentage"`
		Tag        interface{} `json:"tag"`
	} `json:"addressHolders"`
}
