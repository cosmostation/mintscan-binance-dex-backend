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
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/db"
	"github.com/cosmostation/mintscan-binance-dex-backend/mintscan/handlers"

	"github.com/pkg/errors"

	"github.com/gorilla/mux"
)

var (
	// Version is a project's version string.
	Version = "Development"

	// Commit is commit hash of this project.
	Commit = ""
)

func main() {
	l := log.New(os.Stdout, "Mintscan API ", log.Lshortfile|log.LstdFlags)

	// Parse config from configuration file (config.yaml).
	config := config.ParseConfig()

	client := client.NewClient(
		config.Node,
		config.Market,
	)

	db := db.Connect(config.DB)
	err := db.Ping()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to ping database"))
	}

	r := mux.NewRouter()
	r = r.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/account/{address}", handlers.GetAccount).Methods("GET")
	r.HandleFunc("/account/txs/{address}", handlers.GetAccountTxs).Methods("GET")
	r.HandleFunc("/asset", handlers.GetAsset).Methods("GET")
	r.HandleFunc("/assets", handlers.GetAssets).Methods("GET")
	r.HandleFunc("/assets/mini-tokens", handlers.GetAssetsMiniTokens).Methods("GET")
	r.HandleFunc("/assets/txs", handlers.GetAssetTxs).Methods("GET")
	r.HandleFunc("/asset-holders", handlers.GetAssetHolders).Methods("GET")
	r.HandleFunc("/assets-images", handlers.GetAssetsImages).Methods("GET")
	r.HandleFunc("/blocks", handlers.GetBlocks).Methods("GET")
	r.HandleFunc("/fees", handlers.GetFees).Methods("GET")
	r.HandleFunc("/validators", handlers.GetValidators).Methods("GET")
	r.HandleFunc("/validator/{address}", handlers.GetValidator).Methods("GET")
	r.HandleFunc("/market", handlers.GetCoinMarketData).Methods("GET")
	r.HandleFunc("/market/chart", handlers.GetCoinMarketChartData).Methods("GET")
	r.HandleFunc("/orders/{id}", handlers.GetOrders).Methods("GET")
	r.HandleFunc("/stats/assets/chart", handlers.GetAssetsChartHistory).Methods("GET")
	r.HandleFunc("/status", handlers.GetStatus).Methods("GET")
	r.HandleFunc("/tokens", handlers.GetTokens).Methods("GET")
	r.HandleFunc("/txs", handlers.GetTxs).Methods("GET")
	r.HandleFunc("/txs", handlers.GetTxsByTxType).Methods("POST")
	r.HandleFunc("/txs/{hash}", handlers.GetTxByTxHash).Methods("GET")

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // catch-all
		w.Write([]byte("No route is found matching the URL"))
	})

	// Create a new server
	sm := &http.Server{
		Addr:         ":" + config.Web.Port,
		Handler:      handlers.Middleware(r, client, db, l),
		ErrorLog:     l,
		ReadTimeout:  50 * time.Second, // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
	}

	// Start the API server
	go func() {
		l.Printf("Server is running on http://localhost:%s\n", config.Web.Port)
		l.Printf("Version: %s | Commit: %s", Version, Commit)

		err := sm.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	TrapSignal(sm)
}

// TrapSignal traps sigterm or interupt and gracefully shutdown the server
func TrapSignal(sm *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received.
	sig := <-c

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	sm.Shutdown(ctx)

	log.Println("Gracefully shutting down the server: ", sig)
}
