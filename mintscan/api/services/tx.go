package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	var txs []schema.Transaction

	var err error

	switch {
	case before > 0:
		txs, err = db.QueryTxs(limit, before, after, offset)
		if err != nil {
			fmt.Println(err)
		}
	case after > 0:
		txs, err = db.QueryTxs(limit, before, after, offset)
		if err != nil {
			fmt.Println(err)
		}
	case offset >= 0:
		txs, err = db.QueryTxs(limit, before, after, offset)
		if err != nil {
			fmt.Println(err)
		}
	}

	result := make([]*models.ResultTxs, 0)

	for i, tx := range txs {
		msgs := make([]models.Message, 0)
		err := json.Unmarshal([]byte(tx.Messages), &msgs)
		if err != nil {
			fmt.Printf("failed to unmarshal msgs: %t\n", err)
		}

		sigs := make([]models.Signature, 0)
		err = json.Unmarshal([]byte(tx.Signatures), &sigs)
		if err != nil {
			fmt.Printf("failed to unmarshal sigs: %t\n", err)
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
			Timestamp:  tx.Timestamp,
		}

		result = append(result, tempTx)
	}

	utils.Respond(w, result)
	return nil
}
