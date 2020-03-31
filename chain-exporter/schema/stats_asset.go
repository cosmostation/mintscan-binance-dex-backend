package schema

import "time"

// StatAssetInfoList1H defines the schema for asset statistics in an hourly basis
type StatAssetInfoList1H struct {
	ID              int32     `json:"id" sql:",pk"`
	TotalNum        int       `json:"total_num"`
	Name            string    `json:"name" sql:",notnull"`
	Asset           string    `json:"asset" sql:",notnull"`
	Owner           string    `json:"owner" sql:",notnull"`
	Price           float64   `json:"price"`
	Currency        string    `json:"currency"`
	ChangeRange     float64   `json:"change_range"`
	Supply          float64   `json:"supply" sql:",notnull"`
	Marketcap       float64   `json:"marketcap" sql:",notnull"`
	AssetImg        string    `json:"asset_img" sql:",notnull"`
	AssetCreateTime int64     `json:"asset_create_time"`
	Timestamp       time.Time `json:"timestamp" sql:"default:now()"`
}

// StatAssetInfoList24H defines the schema for asset statistics in 24 hourly basis
type StatAssetInfoList24H struct {
	ID              int32     `json:"id" sql:",pk"`
	TotalNum        int       `json:"total_num"`
	Name            string    `json:"name" sql:",notnull"`
	Asset           string    `json:"asset" sql:",notnull"`
	Owner           string    `json:"owner" sql:",notnull"`
	Price           float64   `json:"price"`
	Currency        string    `json:"currency"`
	ChangeRange     float64   `json:"change_range"`
	Supply          float64   `json:"supply" sql:",notnull"`
	Marketcap       float64   `json:"marketcap" sql:",notnull"`
	AssetImg        string    `json:"asset_img" sql:",notnull"`
	AssetCreateTime int64     `json:"asset_create_time"`
	Timestamp       time.Time `json:"timestamp" sql:"default:now()"`
}
