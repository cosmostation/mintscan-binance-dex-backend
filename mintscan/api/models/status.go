package models

import "time"

// Status represents status on the active chain
type Status struct {
	ChainID              string    `json:"chain_id"`
	BlockHeight          int64     `json:"block_height"`
	TotalValidatorNum    int       `json:"total_validator_num"`
	UnjailedValidatorNum int       `json:"unjailed_validator_num"`
	JailedValidatorNum   int       `json:"jailed_validator_num"`
	Timestamp            time.Time `json:"time"`
}
