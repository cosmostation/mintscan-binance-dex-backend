package api

import (
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/codec"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/controllers"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/api/db"

	amino "github.com/tendermint/go-amino"

	"github.com/gorilla/mux"
)

// App wraps up the required variables that are needed in this app
type App struct {
	cdc    *amino.Codec
	client client.Client
	db     *db.Database
	router *mux.Router
}

// NewApp initializes the app with predefined configuration
func NewApp() *App {
	cfg := config.ParseConfig()

	client := client.NewClient(
		cfg.Node.RPCNode,
		cfg.Node.AcceleratedNode,
		cfg.Node.APIServerEndpoint,
		cfg.Market.CoinGeckoEndpoint,
		cfg.Node.ExplorerServerEndpoint,
		cfg.Node.NetworkType,
	)

	app := &App{
		cdc:    codec.Codec,
		client: client,
		db:     db.Connect(cfg.DB),
		router: setRouter(),
	}

	err := app.db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping database "))
	}

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
	controllers.AccountController(a.cdc, a.client, a.db, a.router)
	controllers.AssetController(a.cdc, a.client, a.db, a.router)
	controllers.BlockController(a.cdc, a.client, a.db, a.router)
	controllers.StatusController(a.cdc, a.client, a.db, a.router)
	controllers.StatsController(a.cdc, a.client, a.db, a.router)
	controllers.MarketController(a.cdc, a.client, a.db, a.router)
	controllers.OrderController(a.cdc, a.client, a.db, a.router)
	controllers.TokenController(a.cdc, a.client, a.db, a.router)
	controllers.TxController(a.cdc, a.client, a.db, a.router)
}

// Run runs the API server
func (a *App) Run(port string) {
	log.Print("Server is starting on http://localhost", port, "\n")
	log.Fatal(http.ListenAndServe(port, a.router))
}
