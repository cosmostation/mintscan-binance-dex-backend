package handlers

import (
	"log"
	"net/http"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// Fee is a fee handler
type Fee struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewFee creates a new fee handler with the given params
func NewFee(l *log.Logger, client *client.Client, db *db.Database) *Fee {
	return &Fee{l, client, db}
}

// GetFees returns current fee on the active chain
func (f *Fee) GetFees(rw http.ResponseWriter, r *http.Request) {
	fees, err := f.client.TxMsgFees()
	if err != nil {
		f.l.Printf("failed to fetch tx msg fees: %s", err)
		return
	}

	utils.Respond(rw, fees)
	return
}
