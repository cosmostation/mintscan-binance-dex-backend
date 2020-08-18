package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"

	"github.com/gorilla/mux"
)

// GetAccount returns account information
func GetAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "address is required")
		return
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "address is invalid")
		return
	}

	account, err := s.client.GetAccount(address)
	if err != nil {
		s.l.Printf("failed to request account information: %s\n", err)
	}

	models.Respond(rw, account)
	return
}

// GetAccountTxs returns transactions associated with an account
func GetAccountTxs(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "address is required")
		return
	}

	if len(address) != 42 {
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
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be less than")
		return
	}

	if rows > 50 {
		errors.ErrInvalidParam(rw, http.StatusBadRequest, "'rows' cannot be greater than 50")
		return
	}

	acctTxs, err := s.client.GetAccountTxs(address, page, rows)
	if err != nil {
		s.l.Printf("failed to get account txs: %s\n", err)
	}

	txArray := make([]models.AccountTxArray, 0)

	for _, tx := range acctTxs.TxArray {
		var toAddr string
		if tx.ToAddr != "" {
			toAddr = tx.ToAddr
		}

		tempTxArray := &models.AccountTxArray{
			BlockHeight:   tx.BlockHeight,
			TxHash:        tx.TxHash,
			Code:          tx.Code,
			TxType:        tx.TxType,
			TxAsset:       tx.TxAsset,
			TxQuoteAsset:  tx.TxQuoteAsset,
			Value:         tx.Value,
			TxFee:         tx.TxFee,
			TxAge:         tx.TxAge,
			FromAddr:      tx.FromAddr,
			ToAddr:        toAddr,
			Log:           tx.Log,
			ConfirmBlocks: tx.ConfirmBlocks,
			Memo:          tx.Memo,
			Source:        tx.Source,
			Timestamp:     tx.TimeStamp,
		}

		// txType TRANSFER shouldn't throw message data
		var data models.AccountTxData
		if tx.Data != "" {
			err = json.Unmarshal([]byte(tx.Data), &data)
			if err != nil {
				s.l.Printf("failed to unmarshal AssetTxData: %s", err)
			}

			tempTxArray.Message = &data
		}

		txArray = append(txArray, *tempTxArray)
	}

	result := &models.ResultAccountTxs{
		TxNums:  acctTxs.TxNums,
		TxArray: txArray,
	}

	models.Respond(rw, result)
	return
}
