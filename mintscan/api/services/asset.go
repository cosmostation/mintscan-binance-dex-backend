package services

import (
	"fmt"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
)

// GetAssets returns assets based upon the request params
func GetAssets(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	vals, err := client.Validators()
	fmt.Println("err: ", err)

	for _, val := range vals {
		fmt.Println(val.Description.Moniker)
		fmt.Println(val.AccountAddress)
		fmt.Println(val.OperatorAddress)
	}

	return nil
}
