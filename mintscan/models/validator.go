package models

import (
	"time"
)

type (
	// Validator defines the structure for validator API
	Validator struct {
		OperatorAddress string      `json:"operator_address" sql:",notnull, unique"`
		ConsensusPubKey string      `json:"consensus_pubkey" sql:",notnull, unique"`
		Jailed          bool        `json:"jailed"`
		Status          string      `json:"status"`
		Tokens          string      `json:"tokens"`
		Power           int64       `json:"power"`
		DelegatorShares string      `json:"delegator_shares"`
		Description     Description `json:"description"`
		UnbondingHeight int64       `json:"unbonding_height"`
		UnbondingTime   string      `json:"unbonding_time"`
		Commission      Commission  `json:"commission"`
	}

	// Description wraps validator's description information
	Description struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	}

	// Commission wraps validator's commission information
	Commission struct {
		Rate          string    `json:"rate"`
		MaxRate       string    `json:"max_rate"`
		MaxChangeRate string    `json:"max_change_rate"`
		UpdateTime    time.Time `json:"update_time"`
	}
)
