package models

import "time"

// Status defines the structure for current status on the active chain
type Status struct {
	ChainID           string    `json:"chain_id"`
	BlockTime         float64   `json:"block_time"`
	LatestBlockHeight int64     `json:"latest_block_height"`
	TotalValidatorNum int       `json:"total_validator_num"`
	Timestamp         time.Time `json:"timestamp"`
}
