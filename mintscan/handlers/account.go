package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
)

// GetAccount returns account information
func GetAccount(c *gin.Context) {
	address := c.Params.ByName("address")

	if address == "" {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "address is required")
		return
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "address is invalid")
		return
	}

	account, err := s.client.GetAccount(address)
	if err != nil {
		s.l.Printf("failed to request account information: %s\n", err)
	}

	models.Respond(c.Writer, account)
	return
}

// GetAccountTxs returns transactions associated with an account
func GetAccountTxs(c *gin.Context) {
	q := c.Request.URL.Query()
	address := c.Params.ByName("address")

	if address == "" {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "address is required")
		return
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "address is invalid")
		return
	}

	page := int(1)
	rows := int(10)

	if len(q["page"]) > 0 {
		page, _ = strconv.Atoi(q["page"][0])
	}

	if len(q["rows"]) > 0 {
		rows, _ = strconv.Atoi(q["rows"][0])
	}

	if rows < 1 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be less than")
		return
	}

	if rows > 50 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "'rows' cannot be greater than 50")
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

	models.Respond(c.Writer, result)
	return
}
