package handlers

import (
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// GetFees returns current fee on the active chain
func GetFees(rw http.ResponseWriter, r *http.Request) {
	fees, err := s.client.GetTxMsgFees()
	if err != nil {
		s.l.Printf("failed to fetch tx msg fees: %s", err)
		return
	}

	utils.Respond(rw, fees)
	return
}
