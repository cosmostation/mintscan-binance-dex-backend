package cron

// import (
// 	"fmt"
// 	"log"

// 	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

// 	"github.com/pkg/errors"

// 	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/client"
// 	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/config"
// 	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/db"
// 	// "github.com/robfig/cron"
// )

// // Cron wraps the required params to export blockchain
// type Cron struct {
// 	client client.Client
// 	db     *db.Database
// }

// // NewCron returns Cron
// func NewCron() Cron {
// 	cfg := config.ParseConfig()

// 	client := client.NewClient(
// 		cfg.Node.RPCNode,
// 		cfg.Node.AcceleratedNode,
// 		cfg.Node.APIServerEndpoint,
// 		cfg.Node.ExplorerServerEndpoint,
// 		cfg.Node.NetworkType,
// 	)

// 	db := db.Connect(cfg.DB)

// 	// Ping database to verify connection is succeeded
// 	err := db.Ping()
// 	if err != nil {
// 		log.Fatal(errors.Wrap(err, "failed to ping database."))
// 	}

// 	// Create database tables
// 	db.CreateTables()

// 	return Cron{client, db}
// }

// // Start starts to create cron jobs, which will fetch asset information list and
// // store the data in database every hour and every day
// func (c *Cron) Start() error {
// 	// cron := cron.New()

// 	fmt.Println("CronJob Started")

// 	assetInfoList1H, err := c.getAssetInfoList1H()
// 	if err != nil {
// 		log.Printf("failed to get asset into list 1H: %s", err)
// 	}

// 	assetInfoList24H, err := c.getAssetInfoList24H()
// 	if err != nil {
// 		log.Printf("failed to get asset into list 24H: %s", err)
// 	}

// 	fmt.Println("assetInfoList1H: ", assetInfoList1H)
// 	fmt.Println("assetInfoList24H: ", assetInfoList24H)

// 	// // Every hour
// 	// cron.AddFunc("0 0 * * * *", func() {
// 	// 	c.db.SaveAssetInfoList1H(assetInfoList1H)
// 	// 	log.Println("successfully saved asset information list 1H")
// 	// })

// 	// // Every 24 hours at 1:00 AM UTC timezone
// 	// cron.AddFunc("0 0 1 * * *", func() {
// 	// 	c.db.SaveAssetInfoList24H(assetInfoList24H)
// 	// 	log.Println("successfully saved asset information list 24H")
// 	// })

// 	// go c.Start()

// 	// // Allow graceful closing of the governance loop
// 	// signalCh := make(chan os.Signal, 1)
// 	// signal.Notify(signalCh, os.Interrupt)
// 	// <-signalCh

// 	// Test
// 	c.db.SaveAssetInfoList1H(assetInfoList1H)
// 	c.db.SaveAssetInfoList24H(assetInfoList24H)

// 	return nil
// }

// // getAssetInfoList1H fetches asset information list and return them
// func (c *Cron) getAssetInfoList1H() ([]*schema.StatAssetInfoList1H, error) {
// 	assetInfoList := make([]*schema.StatAssetInfoList1H, 0)

// 	page := int(1)
// 	rows := int(200)

// 	assets, err := c.client.AssetInfoList(page, rows)
// 	if err != nil {
// 		return assetInfoList, err
// 	}

// 	for _, asset := range assets.AssetInfoList {
// 		tempAssetInfoList := &schema.StatAssetInfoList1H{
// 			TotalNum:        assets.TotalNum,
// 			Name:            asset.Name,
// 			Asset:           asset.Asset,
// 			Owner:           asset.Owner,
// 			Price:           asset.Price,
// 			Currency:        asset.QuoteUnit,
// 			ChangeRange:     asset.ChangeRange,
// 			Supply:          asset.Supply,
// 			Marketcap:       asset.Supply * asset.Price,
// 			AssetImg:        asset.AssetImg,
// 			AssetCreateTime: asset.AssetCreateTime,
// 		}

// 		assetInfoList = append(assetInfoList, tempAssetInfoList)
// 	}

// 	return assetInfoList, nil
// }

// // getAssetInfoList24H fetches asset information list and return them
// func (c *Cron) getAssetInfoList24H() ([]*schema.StatAssetInfoList24H, error) {
// 	assetInfoList := make([]*schema.StatAssetInfoList24H, 0)

// 	page := int(1)
// 	rows := int(200)

// 	assets, err := c.client.AssetInfoList(page, rows)
// 	if err != nil {
// 		return assetInfoList, err
// 	}

// 	for _, asset := range assets.AssetInfoList {
// 		tempAssetInfoList := &schema.StatAssetInfoList24H{
// 			TotalNum:        assets.TotalNum,
// 			Name:            asset.Name,
// 			Asset:           asset.Asset,
// 			Owner:           asset.Owner,
// 			Price:           asset.Price,
// 			Currency:        asset.QuoteUnit,
// 			ChangeRange:     asset.ChangeRange,
// 			Supply:          asset.Supply,
// 			Marketcap:       asset.Supply * asset.Price,
// 			AssetImg:        asset.AssetImg,
// 			AssetCreateTime: asset.AssetCreateTime,
// 		}

// 		assetInfoList = append(assetInfoList, tempAssetInfoList)
// 	}

// 	return assetInfoList, nil
// }
