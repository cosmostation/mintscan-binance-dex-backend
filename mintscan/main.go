package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/client"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/config"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/db"
	"github.com/InjectiveLabs/injective-explorer-mintscan-backend/mintscan/handlers"
)

var (
	// Version is a project's version string.
	Version = "dev"

	// Commit is commit hash of this project.
	Commit = ""
)

func main() {
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

	r := gin.Default()

	if Version != "dev" && os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r.Use(cors.Default())
	v1 := r.Group("/v1")
	v1.GET("/account/:address", handlers.GetAccount)
	v1.GET("/accountTxs/:address", handlers.GetAccountTxs)
	v1.GET("/asset", handlers.GetAsset)
	v1.GET("/assets", handlers.GetAssets)
	v1.GET("/assets/mini-tokens", handlers.GetAssetsMiniTokens)
	v1.GET("/assets/txs", handlers.GetAssetTxs)
	v1.GET("/asset-holders", handlers.GetAssetHolders)
	v1.GET("/assets-images", handlers.GetAssetsImages)
	v1.GET("/blocks", handlers.GetBlocks)
	v1.GET("/fees", handlers.GetFees)
	v1.GET("/validators", handlers.GetValidators)
	v1.GET("/validator/:address", handlers.GetValidator)
	v1.GET("/market", handlers.GetCoinMarketData)
	v1.GET("/market/chart", handlers.GetCoinMarketChartData)
	v1.GET("/orders/:id", handlers.GetOrders)
	v1.GET("/stats/assets/chart", handlers.GetAssetsChartHistory)
	v1.GET("/status", handlers.GetStatus)
	v1.GET("/tokens", handlers.GetTokens)
	v1.GET("/txs", handlers.GetTxs)
	v1.POST("/txs", handlers.GetTxsByTxType)
	v1.GET("/txs/:hash", handlers.GetTxByTxHash)

	// Create a new server
	sm := &http.Server{
		Addr:         ":" + config.Web.Port,
		Handler:      handlers.Middleware(r, client, db, log.DefaultLogger),
		ReadTimeout:  50 * time.Second, // max time to read request from the client
		WriteTimeout: 10 * time.Second, // max time to write response to the client
	}

	// Start the API server
	go func() {
		log.Printf("Server is running on http://localhost:%s\n", config.Web.Port)
		log.Printf("Version: %s | Commit: %s", Version, Commit)

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
