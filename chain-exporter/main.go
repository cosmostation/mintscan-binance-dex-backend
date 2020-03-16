package main

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/cron"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/exporter"
)

func main() {
	// Start exporting chain data
	exporter := exporter.NewExporter()
	exporter.Start()

	// Start cron jobs to store data for every certain time
	cron := cron.NewCron()
	cron.Start()
}
