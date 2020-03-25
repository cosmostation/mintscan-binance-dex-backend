package services

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// GetBlocks returns blocks based upon the request params
func GetBlocks(db *db.Database, w http.ResponseWriter, r *http.Request) error {
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
		errors.ErrOverMaxLimit(w, http.StatusUnauthorized)
		return nil
	}

	blocks, err := db.QueryBlocks(before, after, limit)
	if err != nil {
		log.Printf("failed to query blocks: %s\n", err)
	}

	if len(blocks) <= 0 {
		return nil
	}

	result, err := setBlocks(db, blocks)
	if err != nil {
		log.Printf("failed to set blocks: %s\n", err)
	}

	latestBlockHeight, err := db.QueryLatestBlockHeight()
	if err != nil {
		log.Printf("failed to query latest block height: %s\n", err)
	}

	// Handling before and after since their ordering data is different
	if after >= 0 {
		result.Paging.Total = int32(latestBlockHeight)
		result.Paging.Before = int32(result.Data[0].Height)
		result.Paging.After = int32(result.Data[len(result.Data)-1].Height)
	} else {
		result.Paging.Total = int32(latestBlockHeight)
		result.Paging.Before = int32(result.Data[len(result.Data)-1].Height)
		result.Paging.After = int32(result.Data[0].Height)
	}

	utils.Respond(w, result)
	return nil
}

// setBlocks handles blocks and return result response
func setBlocks(db *db.Database, blocks []schema.Block) (*models.ResultBlocks, error) {
	data := make([]models.BlockData, 0)

	for _, block := range blocks {
		resultTxs := make([]models.Txs, 0)

		// Check if any transaction exists in this block
		if block.NumTxs > 0 {
			txs, _ := db.QueryTx(block.Height)
			for _, tx := range txs {
				msgs := make([]models.Message, 0)
				err := json.Unmarshal([]byte(tx.Messages), &msgs)
				if err != nil {
					log.Printf("failed to unmarshal msgs: %s\n", err)
				}

				txResult := true
				if tx.Code != 0 {
					txResult = false
				}

				tempTx := &models.Txs{
					Height:    tx.Height,
					Result:    txResult,
					TxHash:    tx.TxHash,
					Messages:  msgs,
					Memo:      tx.Memo,
					Code:      tx.Code,
					Timestamp: tx.Timestamp,
				}

				resultTxs = append(resultTxs, *tempTx)
			}
		}

		tempData := &models.BlockData{
			Height:        block.Height,
			Proposer:      block.Proposer,
			Moniker:       block.Moniker,
			BlockHash:     block.BlockHash,
			ParentHash:    block.ParentHash,
			NumPrecommits: block.NumPrecommits,
			NumTxs:        block.NumTxs,
			TotalTxs:      block.TotalTxs,
			Txs:           resultTxs,
			Timestamp:     block.Timestamp,
		}

		data = append(data, *tempData)
	}

	result := &models.ResultBlocks{
		Data: data,
	}

	return result, nil
}
