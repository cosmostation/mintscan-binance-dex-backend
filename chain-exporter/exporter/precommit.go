package exporter

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// getPreCommits parses validators information and wrap into Precommit schema struct
func (ex *Exporter) getPreCommits(commit *tmtypes.Commit, vals *tmctypes.ResultValidators) ([]*schema.PreCommit, error) {
	precommits := make([]*schema.PreCommit, 0)

	if len(commit.Precommits) > 0 {
		for _, precommit := range commit.Precommits {
			if precommit != nil { // avoid nil-Vote
				valAddr := precommit.ValidatorAddress.String()

				val := findValidatorByAddr(valAddr, vals)
				if val == nil {
					return nil, fmt.Errorf("failed to find validator by address %s for block %d", valAddr, precommit.Height)
				}

				tempPreCommit := &schema.PreCommit{
					Height:           precommit.Height,
					Round:            precommit.Round,
					ValidatorAddress: valAddr,
					VotingPower:      val.VotingPower,
					ProposerPriority: val.ProposerPriority,
					Timestamp:        precommit.Timestamp,
				}

				precommits = append(precommits, tempPreCommit)
			}
		}
	}

	return precommits, nil
}

// findValidatorByAddr finds a validator by a HEX address given a set of
// Tendermint validators for a particular block. If no validator is found, nil is returned.
func findValidatorByAddr(addrHex string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if addrHex == val.Address.String() {
			return val
		}
	}

	return nil
}
