//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=models/models.cfg.yaml api/openapi.yaml

package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	"github.com/Taraxa-project/taraxa-indexer/api"
	"github.com/Taraxa-project/taraxa-indexer/internal/auth"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/indexer"
	"github.com/Taraxa-project/taraxa-indexer/internal/logging"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	migration "github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble/migrations"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	log "github.com/sirupsen/logrus"
)

var (
	http_port            *int
	metrics_port         *int
	blockchain_ws        *string
	data_dir             *string
	log_level            *string
	sync_queue_limit     *int
	chain_stats_interval *int
	auth_username        *string
	auth_password        *string
	retry_time           time.Duration = 5 * time.Second
)

func init() {
	http_port = flag.Int("http_port", 8080, "port to listen")
	metrics_port = flag.Int("metrics_port", 2112, "metrics http port")
	blockchain_ws = flag.String("blockchain_ws", "wss://ws.mainnet.taraxa.io", "ws url to connect to blockchain")
	data_dir = flag.String("data_dir", "./data", "path to directory where indexer database will be saved")
	log_level = flag.String("log_level", "info", "minimum log level. could be only [trace, debug, info, warn, error, fatal]")
	sync_queue_limit = flag.Int("sync_queue_limit", 10, "limit of blocks in the sync queue")
	chain_stats_interval = flag.Int("chain_stats_interval", 100, "interval for saving chain stats")
	auth_username = flag.String("auth_username", "taraxa", "username for protected endpoints (required if auth_password is set)")
	auth_password = flag.String("auth_password", "taraxa", "password for protected endpoints (required if auth_username is set)")

	flag.Parse()

	logging.Config(filepath.Join(*data_dir, "logs"), *log_level)
	log.Print("\n\n\n")
	log.WithFields(log.Fields{
		"http_port":     *http_port,
		"blockchain_ws": *blockchain_ws,
		"data_dir":      *data_dir,
		"log_level":     *log_level}).
		Info("Application initialized")
}

func setupCloseHandler(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go func() {
		<-c
		fn()
		os.Exit(0)
	}()
}

func connectToChain() (client common.Client, err error) {
	for {
		client, err = chain.NewWsClient(*blockchain_ws)
		if err == nil {
			break
		}
		log.WithError(err).Error("Can't connect to chain")
		time.Sleep(retry_time)
	}
	return
}

func main() {
	log.Info("Starting Taraxa Indexer")
	st := pebble.NewStorage(filepath.Join(*data_dir, "db"))
	setupCloseHandler(func() { _ = st.Close() })
	fin := st.GetFinalizationData()

	manager := migration.NewManager(st)
	err := manager.ApplyAll()
	if err != nil {
		log.WithError(err).Fatal("Error applying migrations")
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		log.WithError(err).Fatal("Error loading swagger spec")
	}
	client, err := connectToChain()
	if err != nil {
		log.WithError(err).Fatal("Error connecting to chain")
	}

	swagger.Servers = nil

	e := echo.New()
	e.Use(echomiddleware.OapiRequestValidator(swagger))
	// Add http error handler to return a proper error JSON on request error
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": err.Error()})
	}

	c := common.DefaultConfig()
	c.SyncQueueLimit = uint64(*sync_queue_limit)
	c.ChainStatsInterval = *chain_stats_interval
	c.AuthUsername = *auth_username
	c.AuthPassword = *auth_password

	// Setup OpenAPI authentication
	auth.SetupOpenAPIAuth(e, swagger, auth.Config{
		Username: c.AuthUsername,
		Password: c.AuthPassword,
	})

	log.WithFields(log.Fields{"pbft_count": fin.PbftCount, "dag_count": fin.DagCount, "trx_count": fin.TrxCount}).Info("Loaded db with")
	chainStats := chain.MakeStats(c.ChainStatsInterval)
	apiHandler := api.NewApiHandler(st, c, chainStats)
	api.RegisterHandlers(e, apiHandler)

	go indexer.MakeAndRun(client, st, c, chainStats, retry_time)

	// start a http server for prometheus on a separate go routine
	go metrics.RunPrometheusServer(":" + strconv.FormatInt(int64(*metrics_port), 10))

	err = e.Start(":" + strconv.FormatInt(int64(*http_port), 10))
	log.WithError(err).Fatal("Can't start http server")
}
