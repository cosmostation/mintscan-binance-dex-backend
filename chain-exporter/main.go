package main

import (
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/exporter"
)

func main() {
	// Start exporting chain data.
	exporter := exporter.NewExporter()
	exporter.Start()
}
