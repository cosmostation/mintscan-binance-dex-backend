package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/pkg/errors"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/controllers"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"

	"github.com/binance-chain/go-sdk/client/rpc"

	amino "github.com/tendermint/go-amino"

	"github.com/gorilla/mux"

	resty "gopkg.in/resty.v1"
)

// App wraps up the required variables that are needed in this app
type App struct {
	cdc       *amino.Codec
	cfg       *config.Config
	db        *db.Database
	router    *mux.Router
	rpcClient rpc.Client
}

// NewApp initializes the app with predefined configuration
func NewApp() *App {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	cfg := config.ParseConfig()

	app := &App{
		cdc:       codec.Codec,
		cfg:       cfg,
		db:        db.Connect(cfg.DB),
		router:    setRouter(),
		rpcClient: rpc.NewRPCClient(cfg.Node.RPCNode, cfg.Node.NetworkType), // Binance Chain RPC Client
	}

	// Ping database to verify connection is succeeded
	err := app.db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping database "))
	}

	// Set timeout for request
	resty.SetTimeout(5 * time.Second)

	// Set controllers
	app.setControllers()

	// Start server
	app.Run(":" + cfg.Web.Port)

	return app
}

// setRouter sets router config
func setRouter() *mux.Router {
	r := mux.NewRouter()
	r = r.PathPrefix("/v1").Subrouter()

	return r
}

// setControllers sets controllers
func (a *App) setControllers() {
	controllers.BlockController(a.cdc, a.cfg, a.db, a.router, a.rpcClient)
	controllers.TxController(a.cdc, a.cfg, a.db, a.router, a.rpcClient)
}

// Run the app
func (a *App) Run(port string) {
	fmt.Print("Server is starting on http://localport", port, "\n")
	log.Fatal(http.ListenAndServe(port, a.router))
}
