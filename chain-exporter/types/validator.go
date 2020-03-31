package types

import (
	"encoding/json"
)

type (
	// Validator defines the structure for validator information
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
		UnbondingTime      string          `json:"unbonding_time"`
		Commission         Commission      `json:"commission"`
	}

	// Description wraps description of a validator
	Description struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	}

	// Commission wrpas general commission information about a validator
	Commission struct {
		Rate          string `json:"rate"`
		MaxRate       string `json:"max_rate"`
		MaxChangeRate string `json:"max_change_rate"`
		UpdateTime    string `json:"update_time"`
	}
)
