package services

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
	"github.com/gorilla/mux"
)

// GetAccount returns account information
func GetAccount(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "address is required")
		return nil
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "address is invalid")
		return nil
	}

	account, err := c.Account(address)
	if err != nil {
		log.Printf("failed to request account information: %s\n", err)
	}

	utils.Respond(w, account)
	return nil
}

// GetAccountTxs returns transactions associated with an account
func GetAccountTxs(c *client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(w, http.StatusBadRequest, "address is required")
		return nil
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "address is invalid")
		return nil
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
		errors.ErrInvalidParam(w, http.StatusBadRequest, "'rows' cannot be less than")
		return nil
	}

	if rows > 50 {
		errors.ErrInvalidParam(w, http.StatusBadRequest, "'rows' cannot be greater than 50")
		return nil
	}

	acctTxs, err := c.AccountTxs(address, page, rows)
	if err != nil {
		log.Printf("failed to get account txs: %s\n", err)
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
				log.Printf("failed to unmarshal AssetTxData: %s", err)
			}

			tempTxArray.Message = &data
		}

		txArray = append(txArray, *tempTxArray)
	}

	result := &models.ResultAccountTxs{
		TxNums:  acctTxs.TxNums,
		TxArray: txArray,
	}

	utils.Respond(w, result)
	return nil
}
