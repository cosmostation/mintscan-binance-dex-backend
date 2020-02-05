package schema

import "time"

// Validator represents validator information
type Validator struct {
	ID               int32     `json:"id" sql:",pk"`
	Moniker          string    `json:"moniker"`
	ValidatorAddress string    `json:"validator_address" sql:",notnull, unique"`
	ConsensusPubKey  string    `json:"consensus_pubkey" sql:",notnull, unique"`
	VotingPower      int64     `json:"voting_power"`
	Timestamp        time.Time `json:"timestamp" sql:"default:now()"`
}
