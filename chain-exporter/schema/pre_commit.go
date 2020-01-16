package schema

import "time"

// PreCommitInfo represents the information about precommit state
type PreCommitInfo struct {
	ID               int32     `json:"id" sql:",pk"`
	Height           int32     `json:"height" sql:",notnull"`
	Round            int32     `json:"round" sql:",notnull"`
	Proposer         string    `json:"proposer" sql:",notnull"`
	VotingPower      int32     `json:"voting_power" sql:",notnull"`
	ProposerPriority int32     `json:"proposer_priority" sql:",notnull"`
	Timestamp        time.Time `json:"timestamp" sql:"default:now()"`
}
