package models

import (
	"encoding/json"
	"time"
)

type (
	// Validator defines the structure for validator API
	Validator struct {
		AccountAddress     string          `json:"account_address" sql:",notnull, unique"`
		OperatorAddress    string          `json:"operator_address" sql:",notnull, unique"`
		ConsensusPubKey    json.RawMessage `json:"consensus_pubkey" sql:",notnull, unique"`
		ConsensusAddress   string          `json:"consensus_address" sql:",notnull, unique"`
		Jailed             bool            `json:"jailed"`
		Status             string          `json:"status"`
		Tokens             string          `json:"tokens"`
		Power              int64           `json:"power"`
		DelegatorShares    string          `json:"delegator_shares"`
		Description        Description     `json:"description"`
		BondHeight         int64           `json:"bond_height"`
		BondIntraTxCounter int64           `json:"bond_intra_tx_counter"`
		UnbondingHeight    int64           `json:"unbonding_height"`
		UnbondingTime      time.Time       `json:"unbonding_time"`
		Commission         Commission      `json:"commission"`
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

// ResultValidators defines the structure for validators response
type ResultValidators struct {
	Moniker          string    `json:"moniker"`
	ValidatorAddress string    `json:"validator_address"`
	ConsensusPubKey  string    `json:"consensus_pubkey"`
	VotingPower      int64     `json:"voting_power"`
	Timestamp        time.Time `json:"timestamp" sql:"default:now()"`
}
