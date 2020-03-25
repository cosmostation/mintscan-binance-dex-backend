package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/controllers"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime | log.Lshortfile)

	if err := run(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to start server."))
	}
}

func run() error {
	cfg := config.ParseConfig()

	client := client.NewClient(
		cfg.Node,
		cfg.Market,
	)

	db := db.Connect(cfg.DB)
	err := db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping database"))
	}

	r := mux.NewRouter()
	r = r.PathPrefix("/v1").Subrouter()

	// set controlelrs
	controllers.AccountController(client, db, r)
	controllers.AssetController(client, db, r)
	controllers.BlockController(client, db, r)
	controllers.StatusController(client, db, r)
	controllers.StatsController(client, db, r)
	controllers.MarketController(client, db, r)
	controllers.OrderController(client, db, r)
	controllers.TokenController(client, db, r)
	controllers.TxController(client, db, r)

	// start the API server
	log.Printf("Server is running on http://localhost:%s\n", cfg.Web.Port)
	if err := http.ListenAndServe(":"+cfg.Web.Port, r); err != nil {
		return fmt.Errorf("http server: %s", err)
	}

	return nil
}
