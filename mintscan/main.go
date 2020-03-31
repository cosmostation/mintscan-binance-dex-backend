package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/client"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/config"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/controllers"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "Mintscan API ", log.LstdFlags)

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

	// create a new server
	sm := &http.Server{
		Addr:         ":" + cfg.Web.Port,
		Handler:      r,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  10 * time.Second,  // max time to read request from the client
		WriteTimeout: 20 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Printf("Server is running on http://localhost:%s\n", cfg.Web.Port)

		err := sm.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	sm.Shutdown(ctx)

	l.Println("Gracefully shutting down the server: ", sig)
}
