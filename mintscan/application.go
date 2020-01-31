package mintscan

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/config"
)

// Parses configuration from config.yaml file and start the server
func main() {
	cfg := config.ParseConfig()

	fmt.Println(cfg)
}
