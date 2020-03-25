package main

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter/cron"
)

func main() {
	cron := cron.NewCron()
	cron.Start()
}
