package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
	"github.com/gorilla/mux"

	cmtypes "github.com/binance-chain/go-sdk/common/types"
)

// Validator is a validator handler
type Validator struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
	nt     cmtypes.ChainNetwork
}

// NewValidator creates a new validator handler with the given params
func NewValidator(l *log.Logger, client *client.Client, db *db.Database, network cmtypes.ChainNetwork) *Validator {
	return &Validator{l, client, db, network}
}

// GetValidators returns validators on the active chain
func (v *Validator) GetValidators(rw http.ResponseWriter, r *http.Request) {
	vals, err := v.db.QueryValidators()
	if err != nil {
		v.l.Printf("failed to query validators: %s", err)
		return
	}

	utils.Respond(rw, vals)
	return
}

// GetValidator returns validator information on the active chain
func (v *Validator) GetValidator(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "address is required")
		return
	}

	switch {
	case strings.HasPrefix(address, v.nt.Bech32ValidatorAddrPrefix()):
		result, err := v.db.QueryValidatorByOperAddr(address)
		if err != nil {
			v.l.Printf("failed to query validator by operator address: %s", err)
			return
		}
		utils.Respond(rw, result)
		return
	case strings.HasPrefix(address, v.nt.Bech32Prefixes()):
		result, err := v.db.QueryValidatorByAccountAddr(address)
		if err != nil {
			v.l.Printf("failed to query validator by account address: %s", err)
			return
		}
		utils.Respond(rw, result)
		return
	case len(address) == 40:
		result, err := v.db.QueryValidatorByConsAddr(address)
		if err != nil {
			v.l.Printf("failed to query validator by consensus address: %s", err)
			return
		}
		utils.Respond(rw, result)
		return
	default:
		result, err := v.db.QueryValidatorByMoniker(address)
		if err != nil {
			v.l.Printf("failed to query validator by moniker: %s", err)
			return
		}
		utils.Respond(rw, result)
		return
	}
}
