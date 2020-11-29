package handlers

import (
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/models"
	"github.com/gin-gonic/gin"
)

// GetFees returns current fee on the active chain
func GetFees(c *gin.Context) {
	fees, err := s.client.GetTxMsgFees()
	if err != nil {
		s.l.Printf("failed to fetch tx msg fees: %s", err)
		return
	}

	models.Respond(c.Writer, fees)
	return
}
