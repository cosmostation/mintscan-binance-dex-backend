package main

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/exporter"
)

func main() {
	exporter := exporter.NewExporter()
	exporter.Start()
}
