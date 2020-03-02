package services

import (
	"fmt"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
)

// GetMarketStats returns current status on the active chain
func GetMarketStats(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("GetMarketStats")

	return nil
}
