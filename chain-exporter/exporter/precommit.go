package exporter

import (
	"fmt"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/schema"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// getPreCommits parses validators information and wrap into Precommit schema struct
func (ex *Exporter) getPreCommits(commit *tmtypes.Commit, vals *tmctypes.ResultValidators) (precommits []*schema.PreCommit, err error) {
	if commit == nil || len(commit.Signatures) == 0 {
		return []*schema.PreCommit{}, nil
	}

	for _, commitSig := range commit.Signatures {
		if commitSig.Absent() {
			continue // OK, some precommits can be missing.
		}

		valAddr := sdk.ConsAddress(commitSig.ValidatorAddress).String()

		val := findValidatorByBechAddr(valAddr, vals)
		if val == nil {
			return nil, fmt.Errorf("failed to find validator by address %s for block %d", valAddr, commit.Height)
		}

		pc := schema.NewPrecommit(schema.PreCommit{
			Height:           commit.Height,
			Round:            commit.Round,
			ValidatorAddress: valAddr,
			VotingPower:      val.VotingPower,
			ProposerPriority: val.ProposerPriority,
			Timestamp:        commitSig.Timestamp,
		})

		precommits = append(precommits, pc)
	}

	return precommits, nil
}

// findValidatorByBechAddr finds a validator by a Bech32 address given a set of
// Tendermint validators for a particular block. If no validator is found, nil is returned.
func findValidatorByBechAddr(addrBech string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if addrBech == sdk.ConsAddress(val.Address).String() {
			return val
		}
	}

	return nil
}
