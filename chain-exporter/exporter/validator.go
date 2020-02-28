package exporter

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"
)

// getValidators parses validators information and wrap into Precommit schema struct
func (ex *Exporter) getValidators(vals []*types.Validator) ([]*schema.Validator, error) {
	validators := make([]*schema.Validator, 0)

	// Looping through validators and insert them if not already exists in database
	for _, val := range vals {
		ok, err := ex.db.ExistValidator(val.ConsensusAddress)
		if !ok {
			tempVal := &schema.Validator{
				Moniker:                 val.Description.Moniker,
				AccountAddress:          val.AccountAddress,
				OperatorAddress:         val.OperatorAddress,
				ConsensusAddress:        val.ConsensusAddress,
				Jailed:                  val.Jailed,
				Status:                  val.Status,
				Tokens:                  val.Tokens,
				VotingPower:             val.Power,
				DelegatorShares:         val.DelegatorShares,
				BondHeight:              val.BondHeight,
				BondIntraTxCounter:      val.BondIntraTxCounter,
				UnbondingHeight:         val.UnbondingHeight,
				UnbondingTime:           val.UnbondingTime,
				CommissionRate:          val.Commission.Rate,
				CommissionMaxRate:       val.Commission.MaxRate,
				CommissionMaxChangeRate: val.Commission.MaxChangeRate,
				CommissionUpdateTime:    val.Commission.UpdateTime,
			}

			validators = append(validators, tempVal)
		}

		if err != nil {
			return nil, fmt.Errorf("unexpected error when checking validator existence: %t", err)
		}
	}

	return validators, nil
}
