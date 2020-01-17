package schema

import "time"

// PreCommitInfo represents the information about precommit state
type PreCommitInfo struct {
	ID               int32     `json:"id" sql:",pk"`
	Height           int64     `json:"height" sql:",notnull"`
	Round            int64     `json:"round" sql:",notnull"`
	Proposer         string    `json:"proposer" sql:",notnull"`
	VotingPower      int64     `json:"voting_power" sql:",notnull"`
	ProposerPriority int64     `json:"proposer_priority" sql:",notnull"`
	Timestamp        time.Time `json:"timestamp" sql:"default:now()"`
}
