package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/errors"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/utils"
)

// Token is a token handler
type Token struct {
	l      *log.Logger
	client *client.Client
	db     *db.Database
}

// NewToken creates a new token handler with the given params
func NewToken(l *log.Logger, client *client.Client, db *db.Database) *Token {
	return &Token{l, client, db}
}

// GetTokens returns assets based upon the request params
func (t *Token) GetTokens(rw http.ResponseWriter, r *http.Request) {
	limit := 100
	offset := 0

	if len(r.URL.Query()["limit"]) > 0 {
		limit, _ = strconv.Atoi(r.URL.Query()["limit"][0])
	}

	if len(r.URL.Query()["offset"]) > 0 {
		offset, _ = strconv.Atoi(r.URL.Query()["offset"][0])
	}

	if limit > 1000 {
		errors.ErrOverMaxLimit(rw, http.StatusUnauthorized)
		return
	}

	tks, _ := t.client.Tokens(limit, offset)

	utils.Respond(rw, tks)
	return
}
