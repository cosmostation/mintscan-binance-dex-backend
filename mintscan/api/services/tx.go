package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/utils"
)

// GetTxs returns transactions based upon the request params
func GetTxs(db *db.Database, w http.ResponseWriter, r *http.Request) error {
	limit := int(100)
	before := int(-1)
	after := int(-1)
	offset := int(0)

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["before"]) > 0 {
		before, _ = strconv.Atoi(r.URL.Query()["before"][0])
	}

	if len(r.URL.Query()["after"]) > 0 {
		after, _ = strconv.Atoi(r.URL.Query()["after"][0])
	}

	if len(r.URL.Query()["offset"]) > 0 {
		offset, _ = strconv.Atoi(r.URL.Query()["offset"][0])
	}

	if limit > 100 {
		errors.ErrOverMaxLimit(w, http.StatusUnauthorized)
		return nil
	}

	latestBlockHeight, err := db.QueryLatestBlockHeight()
	if err != nil {
		fmt.Printf("failed to query latest block height saved in database: %t\n", err)
	}

	totalTxsNum, err := db.QueryTotalTxsNum(latestBlockHeight)
	if err != nil {
		fmt.Printf("failed to query total number of txs: %t\n", err)
	}

	txs, err := db.QueryTxs(limit, before, after, offset)
	if err != nil {
		fmt.Printf("failed to query txs due to: %t\n", err)
	}

	result, err := setTxs(txs, totalTxsNum)
	if err != nil {
		fmt.Printf("failed to set txs: %t\n", err)
	}

	utils.Respond(w, result)
	return nil
}

// GetTxsByType returns transactions based upon the request params
func GetTxsByType(client client.Client, db *db.Database, w http.ResponseWriter, r *http.Request) error {
	limit := int(100) // default limit is 50
	before := int(0)

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["before"]) > 0 {
		before, _ = strconv.Atoi(r.URL.Query()["before"][0])
	}

	if limit > 100 {
		errors.ErrOverMaxLimit(w, http.StatusUnauthorized)
		return nil
	}

	var txrp models.TxRequestPayload
	err := json.NewDecoder(r.Body).Decode(&txrp)
	if err != nil {
		fmt.Printf("failed to decode txrp: %t\n", err)
	}

	// Set the first block time if StartTime is not parsed
	// 2019-04-18 06:07:02.15434+00, which is 1555567622 in unix time
	if txrp.StartTime == 0 {
		txrp.StartTime = 1555567622
	}

	// Set current unix time if EndTime is not parsed
	if txrp.EndTime == 0 {
		txrp.EndTime = time.Now().Unix()
	}

	// Validate transaction message type
	ok := models.ValidatorMsgType(txrp.TxType)
	if !ok {
		errors.ErrInvalidMessageType(w, http.StatusUnauthorized)
		return nil
	}

	latestBlockHeight, err := db.QueryLatestBlockHeight()
	if err != nil {
		fmt.Printf("failed to query latest block height saved in database: %t\n", err)
	}

	totalTxsNum, err := db.QueryTotalTxsNum(latestBlockHeight)
	if err != nil {
		fmt.Printf("failed to query total number of txs: %t\n", err)
	}

	txs, err := db.QueryTxsByType(txrp.TxType, txrp.StartTime, txrp.EndTime, limit, before)
	if err != nil {
		fmt.Printf("failed to query txs due to: %t\n", err)
	}

	result, err := setTxs(txs, totalTxsNum)
	if err != nil {
		fmt.Printf("failed to set txs: %t\n", err)
	}

	utils.Respond(w, result)
	return nil
}

func setTxs(txs []schema.Transaction, totalTxsNum int64) ([]*models.ResultTxs, error) {
	result := make([]*models.ResultTxs, 0)

	for i, tx := range txs {
		msgs := make([]models.Message, 0)
		err := json.Unmarshal([]byte(tx.Messages), &msgs)
		if err != nil {
			return result, fmt.Errorf("failed to unmarshal msgs: %t", err)
		}

		sigs := make([]models.Signature, 0)
		err = json.Unmarshal([]byte(tx.Signatures), &sigs)
		if err != nil {
			return result, fmt.Errorf("failed to unmarshal sigs: %t", err)
		}

		txResult := true
		if tx.Code != 0 {
			txResult = false
		}

		tempTx := &models.ResultTxs{
			ID:         i + 1,
			Height:     tx.Height,
			Result:     txResult,
			TxHash:     tx.TxHash,
			Messages:   msgs,
			Signatures: sigs,
			Memo:       tx.Memo,
			Code:       tx.Code,
			TotalTxs:   totalTxsNum,
			Timestamp:  tx.Timestamp,
		}

		result = append(result, tempTx)
	}

	return result, nil
}
