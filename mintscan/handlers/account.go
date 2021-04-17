package handlers

import (
	"net/http"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/errors"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAccount returns account information
func GetAccount(c *gin.Context) {
	address := c.Params.ByName("address")

	if address == "" {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "address is required")
		return
	}

	if len(address) != 42 {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "address is invalid")
		return
	}

	account, err := s.client.GetAccount(address)
	if err != nil {
		s.l.Printf("failed to request account information: %s\n", err)
	}

	models.Respond(c.Writer, account)
	return
}

// GetAccountTxs returns transactions associated with an account
func GetAccountTxs(c *gin.Context) {
	address := c.Params.ByName("address")

	if address == "" {
		errors.ErrRequiredParam(c.Writer, http.StatusBadRequest, "address is required")
		return
	}

	accAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		errors.ErrInvalidParam(c.Writer, http.StatusBadRequest, "address is invalid")
		return
	}

	accTxs, err := s.db.QueryTxsBySigner(accAddr, 100000)
	if err != nil {
		s.l.Printf("failed to get account txs: %s\n", err)
	}

	result, err := setTxs(accTxs)
	if err != nil {
		s.l.Printf("failed to map account txs: %s\n", err)
	}

	models.Respond(c.Writer, result)
	return
}
