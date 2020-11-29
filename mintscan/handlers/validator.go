package handlers

import (
	"net/http"
	"strings"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"

	ctypes "github.com/InjectiveLabs/sdk-go/chain/types"
)

// GetValidators returns validators on the active chain
func GetValidators(c *gin.Context) {
	vals, err := s.db.QueryValidators()
	if err != nil {
		s.l.Printf("failed to query validators: %s", err)
		return
	}

	models.Respond(c.Writer, vals)
	return
}

// GetValidator returns validator information on the active chain
func GetValidator(c *gin.Context) {
	address := c.Params.ByName("address")
	if address == "" {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "address is required")
		return
	}

	switch {
	case strings.HasPrefix(address, ctypes.Bech32PrefixValAddr):
		result, err := s.db.QueryValidatorByOperAddr(address)
		if err != nil {
			s.l.Printf("failed to query validator by operator address: %s", err)
			return
		}
		models.Respond(c.Writer, result)
		return
	case strings.HasPrefix(address, ctypes.Bech32PrefixAccAddr):
		result, err := s.db.QueryValidatorByAccountAddr(address)
		if err != nil {
			s.l.Printf("failed to query validator by account address: %s", err)
			return
		}
		models.Respond(c.Writer, result)
		return
	case strings.HasPrefix(address, ctypes.Bech32PrefixConsAddr):
		result, err := s.db.QueryValidatorByConsAddr(address)
		if err != nil {
			s.l.Printf("failed to query validator by consensus address: %s", err)
			return
		}
		models.Respond(c.Writer, result)
		return
	default:
		result, err := s.db.QueryValidatorByMoniker(address)
		if err != nil {
			s.l.Printf("failed to query validator by moniker: %s", err)
			return
		}
		models.Respond(c.Writer, result)
		return
	}
}
