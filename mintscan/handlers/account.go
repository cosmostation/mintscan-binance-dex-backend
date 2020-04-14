package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"

	"github.com/tendermint/tmlibs/bech32"

	"github.com/gorilla/mux"
)

// Account is a account handler
type Account struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewAccount creates a new account handler with the given params
func NewAccount(l *log.Logger, client *client.Client, db *db.Database) *Account {
	return &Account{l, client, db}
}

// GetAccount returns account information
func (a *Account) GetAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	addrPrefix, decodedAddress, err := bech32.DecodeAndConvert(address)
	if err != nil {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "address is invalid")
		return
	}

	encodedAddress, err := bech32.ConvertAndEncode(addrPrefix, decodedAddress)
	if err != nil {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "address is invalid")
		return
	}

	account, err := a.client.Account(encodedAddress)
	if err != nil {
		a.l.Printf("failed to request account information: %s\n", err)
	}

	utils.Respond(rw, account)
	return
}

// GetAccountTxs returns transactions associated with an account
func (a *Account) GetAccountTxs(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	addrPrefix, decodedAddress, err := bech32.DecodeAndConvert(address)
	if err != nil {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "address is invalid")
		return
	}

	encodedAddress, err := bech32.ConvertAndEncode(addrPrefix, decodedAddress)
	if err != nil {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "address is invalid")
		return
	}

	page := int(1)
	rows := int(10)

	if len(r.URL.Query()["page"]) > 0 {
		page, _ = strconv.Atoi(r.URL.Query()["page"][0])
	}

	if len(r.URL.Query()["rows"]) > 0 {
		rows, _ = strconv.Atoi(r.URL.Query()["rows"][0])
	}

	if rows < 1 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than 1")
		return
	}

	if rows > 30 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 30")
		return
	}

	acctTxs, err := a.db.QueryAccountTxs(encodedAddress, page, rows)
	if err != nil {
		a.l.Printf("failed to get account txs: %s\n", err)
	}

	result := make([]models.TxData, 0)

	for _, tx := range acctTxs {
		msgs := make([]models.Message, 0)
		_ = json.Unmarshal([]byte(tx.Messages), &msgs)

		txResult := true
		if tx.Code != 0 {
			txResult = false
		}

		tempResult := &models.TxData{
			Height:    tx.Height,
			Result:    txResult,
			TxHash:    tx.TxHash,
			Code:      tx.Code,
			Messages:  msgs,
			Memo:      tx.Memo,
			Timestamp: tx.Timestamp,
		}

		result = append(result, *tempResult)
	}

	utils.Respond(rw, result)
	return
}
