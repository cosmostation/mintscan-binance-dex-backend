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
	s := r.PathPrefix("/v1").Subrouter()

	controllers.AccountController(client, db, s)
	controllers.AssetController(client, db, s)
	controllers.BlockController(client, db, s)
	controllers.StatusController(client, db, s)
	controllers.StatsController(client, db, s)
	controllers.MarketController(client, db, s)
	controllers.OrderController(client, db, s)
	controllers.TokenController(client, db, s)
	controllers.TxController(client, db, s)

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // catch-all
		w.Write([]byte("No route is found matching the URL"))
	})

	// start the API server
	log.Printf("Server is running on http://localhost:%s\n", cfg.Web.Port)
	if err := http.ListenAndServe(":"+cfg.Web.Port, r); err != nil {
		return fmt.Errorf("http server: %s", err)
	}

	return nil
}
