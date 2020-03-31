package schema

import "time"

// PreCommit defines the schema for precommit state information
type PreCommit struct {
	ID               int32     `json:"id" sql:",pk"`
	Height           int64     `json:"height" sql:",notnull"`
	Round            int       `json:"round" sql:",notnull"`
	ValidatorAddress string    `json:"validator_address" sql:",notnull"`
	VotingPower      int64     `json:"voting_power" sql:",notnull"`
	ProposerPriority int64     `json:"proposer_priority" sql:",notnull"`
	Timestamp        time.Time `json:"timestamp" sql:"default:now()"`
}
