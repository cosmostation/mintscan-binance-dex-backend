package schema

import "time"

// Validator defines the schema for validator information
type Validator struct {
	ID                      int32     `json:"id" sql:",pk"`
	Moniker                 string    `json:"moniker"`
	OperatorAddress         string    `json:"operator_address" sql:",notnull, unique"`
	ConsensusAddress        string    `json:"consensus_address" sql:",notnull, unique"`
	Jailed                  bool      `json:"jailed"`
	Status                  string    `json:"status"`
	Tokens                  string    `json:"tokens"`
	VotingPower             int64     `json:"voting_power"`
	DelegatorShares         string    `json:"delegator_shares"`
	UnbondingHeight         int64     `json:"unbonding_height" sql:"default:0"`
	UnbondingTime           string    `json:"unbonding_time"`
	CommissionRate          string    `json:"commission_rate"`
	CommissionMaxRate       string    `json:"commission_max_rate"`
	CommissionMaxChangeRate string    `json:"commission_max_change_rate"`
	CommissionUpdateTime    time.Time `json:"commission_update_time"`
	Timestamp               time.Time `json:"timestamp" sql:"default:now()"`
}
