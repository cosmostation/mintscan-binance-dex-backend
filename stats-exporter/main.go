package main

import (
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/stats-exporter/cron"
)

func main() {
	cron := cron.NewCron()
	cron.Start()
}
