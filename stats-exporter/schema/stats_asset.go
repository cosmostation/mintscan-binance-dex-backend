package schema

import "time"

// StatAssetInfoList1H represents statistics for assets in an hour
type StatAssetInfoList1H struct {
	ID           int32     `json:"id" sql:",pk"`
	Name         string    `json:"name" sql:",notnull"`
	Asset        string    `json:"asset" sql:",notnull"`
	MappedAsset  string    `json:"mapped_asset" sql:",notnull"`
	Price        float64   `json:"price"`
	QuoteUnit    string    `json:"quote_unit"`
	ChangeRange  float64   `json:"change_range"`
	Supply       float64   `json:"supply" sql:",notnull"`
	Marketcap    float64   `json:"marketcap" sql:",notnull"`
	Owner        string    `json:"owner" sql:",notnull"`
	Transactions int       `json:"transactions" sql:",notnull"`
	Holders      int       `json:"holders" sql:",notnull"`
	AssetImage   string    `json:"asset_img"`
	Timestamp    time.Time `json:"timestamp" sql:"default:now()"`
}

// StatAssetInfoList24H represents statistics for assets in 24 hours
type StatAssetInfoList24H struct {
	ID           int32     `json:"id" sql:",pk"`
	Asset        string    `json:"asset" sql:",notnull"`
	MappedAsset  string    `json:"mapped_asset" sql:",notnull"`
	Name         string    `json:"name" sql:",notnull"`
	Price        float64   `json:"price"`
	QuoteUnit    string    `json:"quote_unit"`
	ChangeRange  float64   `json:"change_range"`
	Supply       float64   `json:"supply" sql:",notnull"`
	Marketcap    float64   `json:"marketcap" sql:",notnull"`
	Owner        string    `json:"owner" sql:",notnull"`
	Transactions int       `json:"transactions" sql:",notnull"`
	Holders      int       `json:"holders" sql:",notnull"`
	AssetImage   string    `json:"asset_img"`
	Timestamp    time.Time `json:"timestamp" sql:"default:now()"`
}
