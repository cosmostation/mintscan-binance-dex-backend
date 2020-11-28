package client

import (
	"encoding/json"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/stats-exporter/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/stats-exporter/models"
)

type Client struct {
	exchangeClient *resty.Client
}

// NewClient creates a new client with the given config
func NewClient(cfg config.NodeConfig) *Client {
	exchangeClient := resty.New().
		SetHostURL(cfg.ExchangeAPIEndpoint).
		SetTimeout(time.Duration(10 * time.Second))

	return &Client{
		exchangeClient,
	}
}

// Asset returns particular asset information given an asset name
func (c Client) Asset(assetName string) (models.Asset, error) {
	resp, err := c.exchangeClient.R().Get("/asset?asset=" + assetName)
	if err != nil {
		return models.Asset{}, err
	}

	var asset models.Asset
	err = json.Unmarshal(resp.Body(), &asset)
	if err != nil {
		return models.Asset{}, err
	}

	return asset, nil
}

// Assets returns information of all assets existing in an active chain based upon params
func (c Client) Assets(page int, rows int) (models.Assets, error) {
	resp, err := c.exchangeClient.R().Get("/assets?page=" + strconv.Itoa(page) + "&rows=" + strconv.Itoa(rows))
	if err != nil {
		return models.Assets{}, err
	}

	var assets models.Assets
	err = json.Unmarshal(resp.Body(), &assets)
	if err != nil {
		return models.Assets{}, err
	}

	return assets, nil
}
