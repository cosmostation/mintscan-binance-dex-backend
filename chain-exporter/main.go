package main

import (
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/exporter"
)

func main() {
	// Start exporting chain data.
	exporter := exporter.NewExporter()
	exporter.Start()
}
