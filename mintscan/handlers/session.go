package handlers

import (
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/client"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/db"
)

// Sessions is shorten for s will be used throughout this handler pakcage.
var s *Session

// Session is struct for wrapping both client and db structs.
type Session struct {
	client *client.Client
	db     *db.Database
	l      log.Logger
}

// SetSession set Session object.
func SetSession(client *client.Client, db *db.Database, log log.Logger) {
	s = &Session{client, db, log}
}
