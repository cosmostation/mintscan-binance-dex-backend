package exporter

import (
	"fmt"

	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/schema"
	"github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter/types"

	cmtypes "github.com/binance-chain/go-sdk/common/types"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// getValidators parses validators information and wrap into Precommit schema struct
func (ex *Exporter) getValidators(block *tmctypes.ResultBlock, vals *tmctypes.ResultValidators) ([]*schema.Validator, error) {
	validators := make([]*schema.Validator, 0)

	if len(vals.Validators) > 0 {
		for _, val := range vals.Validators {
			valAddr := val.Address.String()

			// Insert validator only if not exist
			ok, err := ex.db.ExistValidator(valAddr)
			if !ok {
				consPubKey, err := cmtypes.Bech32ifyConsPub(val.PubKey)
				if err != nil {
					return nil, err
				}

				moniker := types.GetValidatorMoniker(valAddr)

				tempVals := &schema.Validator{
					ValidatorAddress: valAddr,
					Moniker:          moniker,
					ConsensusPubKey:  consPubKey,
					VotingPower:      val.VotingPower,
					Timestamp:        block.Block.Time,
				}

				validators = append(validators, tempVals)
			}

			if err != nil {
				return nil, fmt.Errorf("unexpected error when querying validator existence: %t", err)
			}
		}
	}

	return validators, nil
}
