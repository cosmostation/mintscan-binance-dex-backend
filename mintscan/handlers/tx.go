package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/schema"

	"github.com/gorilla/mux"
)

// GetTxs returns transactions based upon the request params
func GetTxs(rw http.ResponseWriter, r *http.Request) {
	before := int(0)
	after := int(-1)
	limit := int(100)

	if len(r.URL.Query()["before"]) > 0 {
		before, _ = strconv.Atoi(r.URL.Query()["before"][0])
	}

	if len(r.URL.Query()["after"]) > 0 {
		after, _ = strconv.Atoi(r.URL.Query()["after"][0])
	}

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if limit > 100 {
		errors.ErrOverMaxLimit(rw, http.StatusUnauthorized)
		return
	}

	txs, err := s.db.QueryTxs(before, after, limit)
	if err != nil {
		s.l.Printf("failed to query txs: %s\n", err)
	}

	if len(txs) <= 0 {
		models.Respond(rw, models.ResultTxs{})
		return
	}

	result, err := setTxs(txs)
	if err != nil {
		s.l.Printf("failed to set txs: %s\n", err)
	}

	totalTxsNum, err := s.db.CountTotalTxsNum()
	if err != nil {
		s.l.Printf("failed to query total number of txs: %s\n", err)
	}

	// Handling before and after since their ordering data is different
	if after >= 0 {
		result.Paging.Total = totalTxsNum
		result.Paging.Before = result.Data[0].ID
		result.Paging.After = result.Data[len(result.Data)-1].ID
	} else {
		result.Paging.Total = totalTxsNum
		result.Paging.Before = result.Data[len(result.Data)-1].ID
		result.Paging.After = result.Data[0].ID
	}

	models.Respond(rw, result)
	return
}

// GetTxByTxHash returns certain transaction information by its tx hash
func GetTxByTxHash(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	tx, err := s.db.QueryTxByHash(hash)
	if err != nil {
		s.l.Printf("failed to query tx: %s\n", err)
		models.Respond(rw, models.TxData{})
		return
	}

	result, err := setTx(tx)
	if err != nil {
		s.l.Printf("failed to set tx: %s\n", err)
	}

	models.Respond(rw, result)
	return
}

// GetTxsByTxType returns transactions based upon the request params
func GetTxsByTxType(rw http.ResponseWriter, r *http.Request) {
	before := int(0)
	after := int(-1)
	limit := int(100)

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["before"]) > 0 {
		before, _ = strconv.Atoi(r.URL.Query()["before"][0])
	}

	if len(r.URL.Query()["after"]) > 0 {
		after, _ = strconv.Atoi(r.URL.Query()["after"][0])
	}

	if limit > 100 {
		errors.ErrOverMaxLimit(rw, http.StatusUnauthorized)
		return
	}

	var txrp models.TxRequestPayload
	err := json.NewDecoder(r.Body).Decode(&txrp)
	if err != nil {
		s.l.Printf("failed to decode txrp: %s\n", err)
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
		errors.ErrInvalidMessageType(rw, http.StatusUnauthorized)
		return
	}

	txs, err := s.db.QueryTxsByType(txrp.TxType, txrp.StartTime, txrp.EndTime, before, after, limit)
	if err != nil {
		s.l.Printf("failed to query txs: %s\n", err)
	}

	if len(txs) <= 0 {
		return
	}

	result, err := setTxs(txs)
	if err != nil {
		s.l.Printf("failed to set txs: %s\n", err)
	}

	totalTxsNum, err := s.db.CountTotalTxsNum()
	if err != nil {
		s.l.Printf("failed to query total number of txs: %s\n", err)
	}

	// Handling before and after since their ordering data is different
	if after >= 0 {
		result.Paging.Total = totalTxsNum
		result.Paging.Before = result.Data[0].ID
		result.Paging.After = result.Data[len(result.Data)-1].ID
	} else {
		result.Paging.Total = totalTxsNum
		result.Paging.Before = result.Data[len(result.Data)-1].ID
		result.Paging.After = result.Data[0].ID
	}

	models.Respond(rw, result)
	return
}

// setTx handles tx and return result response
func setTx(tx schema.Transaction) (*models.TxData, error) {
	msgs := make([]models.Message, 0)
	err := json.Unmarshal([]byte(tx.Messages), &msgs)
	if err != nil {
		return &models.TxData{}, fmt.Errorf("failed to unmarshal msgs: %s", err)
	}

	sigs := make([]models.Signature, 0)
	err = json.Unmarshal([]byte(tx.Signatures), &sigs)
	if err != nil {
		return &models.TxData{}, fmt.Errorf("failed to unmarshal sigs: %s", err)
	}

	txResult := true
	if tx.Code != 0 {
		txResult = false
	}

	result := &models.TxData{
		Height:     tx.Height,
		Result:     txResult,
		TxHash:     tx.TxHash,
		Messages:   msgs,
		Signatures: sigs,
		Memo:       tx.Memo,
		Code:       tx.Code,
		Timestamp:  tx.Timestamp,
	}

	return result, nil
}

// setTxs handles txs and return result response
func setTxs(txs []schema.Transaction) (*models.ResultTxs, error) {
	data := make([]models.TxData, 0)

	for _, tx := range txs {
		msgs := make([]models.Message, 0)
		err := json.Unmarshal([]byte(tx.Messages), &msgs)
		if err != nil {
			return &models.ResultTxs{}, fmt.Errorf("failed to unmarshal msgs: %s", err)
		}

		sigs := make([]models.Signature, 0)
		err = json.Unmarshal([]byte(tx.Signatures), &sigs)
		if err != nil {
			return &models.ResultTxs{}, fmt.Errorf("failed to unmarshal sigs: %s", err)
		}

		txResult := true
		if tx.Code != 0 {
			txResult = false
		}

		tempData := &models.TxData{
			ID:         tx.ID,
			Height:     tx.Height,
			Result:     txResult,
			TxHash:     tx.TxHash,
			Messages:   msgs,
			Signatures: sigs,
			Memo:       tx.Memo,
			Code:       tx.Code,
			Timestamp:  tx.Timestamp,
		}

		data = append(data, *tempData)
	}

	result := &models.ResultTxs{
		Data: data,
	}

	return result, nil
}
