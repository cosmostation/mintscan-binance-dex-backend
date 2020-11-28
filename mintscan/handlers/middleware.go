package handlers

import (
	"github.com/tomasen/realip"
	log "github.com/xlab/suplog"
	"net/http"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/client"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/db"
)

// TODO: Response Status Code needs to be logged. Find out how others do to resolve this, maybe third party library?
// Should we use key:value pairs for better json formatted output?

// Middleware logs incoming requests and calls next handler
func Middleware(next http.Handler, c *client.Client, db *db.Database, l log.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		clientIP := realip.FromRequest(r)
		l.Printf("%s %s [%s]", r.Method, r.URL, clientIP)

		// Session will wrap both client and database and be used throughout all handlers.
		SetSession(c, db, l)

		next.ServeHTTP(rw, r)
	})
}
