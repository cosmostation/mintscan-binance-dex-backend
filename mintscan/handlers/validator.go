package handlers

import (
	"net/http"
	"strings"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/models"
	"github.com/gorilla/mux"

	cmtypes "github.com/binance-chain/go-sdk/common/types"
)

// GetValidators returns validators on the active chain
func GetValidators(rw http.ResponseWriter, r *http.Request) {
	vals, err := s.db.QueryValidators()
	if err != nil {
		s.l.Printf("failed to query validators: %s", err)
		return
	}

	models.Respond(rw, vals)
	return
}

// GetValidator returns validator information on the active chain
func GetValidator(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	if address == "" {
		errors.ErrRequiredParam(rw, http.StatusBadRequest, "address is required")
		return
	}

	switch {
	case strings.HasPrefix(address, cmtypes.Network.Bech32ValidatorAddrPrefix()):
		result, err := s.db.QueryValidatorByOperAddr(address)
		if err != nil {
			s.l.Printf("failed to query validator by operator address: %s", err)
			return
		}
		models.Respond(rw, result)
		return
	case strings.HasPrefix(address, cmtypes.Network.Bech32Prefixes()):
		result, err := s.db.QueryValidatorByAccountAddr(address)
		if err != nil {
			s.l.Printf("failed to query validator by account address: %s", err)
			return
		}
		models.Respond(rw, result)
		return
	case len(address) == 40:
		result, err := s.db.QueryValidatorByConsAddr(address)
		if err != nil {
			s.l.Printf("failed to query validator by consensus address: %s", err)
			return
		}
		models.Respond(rw, result)
		return
	default:
		result, err := s.db.QueryValidatorByMoniker(address)
		if err != nil {
			s.l.Printf("failed to query validator by moniker: %s", err)
			return
		}
		models.Respond(rw, result)
		return
	}
}
