package exporter

import (
	"fmt"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/schema"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/chain-exporter/types"
)

// getValidators parses validators information and wrap into Precommit schema struct
func (ex *Exporter) getValidators(vals []*types.Validator) (validators []*schema.Validator, err error) {
	for _, val := range vals {
		ok, err := ex.db.ExistValidator(val.ConsensusPubKey)
		if !ok {
			val := &schema.Validator{
				Moniker:                 val.Description.Moniker,
				OperatorAddress:         val.OperatorAddress,
				Jailed:                  val.Jailed,
				Status:                  val.Status,
				Tokens:                  val.Tokens,
				VotingPower:             val.Power,
				DelegatorShares:         val.DelegatorShares,
				UnbondingHeight:         val.UnbondingHeight,
				UnbondingTime:           val.UnbondingTime,
				CommissionRate:          val.Commission.Rate,
				CommissionMaxRate:       val.Commission.MaxRate,
				CommissionMaxChangeRate: val.Commission.MaxChangeRate,
				CommissionUpdateTime:    val.Commission.UpdateTime,
			}

			validators = append(validators, val)
		}

		if err != nil {
			return nil, fmt.Errorf("unexpected error when checking validator existence: %s", err)
		}
	}

	return validators, nil
}
