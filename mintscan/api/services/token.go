package services

import (
	"fmt"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
)

// GetTokens returns assets based upon the request params
func GetTokens(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("GetTokens")
	return nil
}
