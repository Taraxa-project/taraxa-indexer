//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=models/models.cfg.yaml api/openapi.yaml

package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/api"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/indexer"
	"github.com/Taraxa-project/taraxa-indexer/internal/lara"
	"github.com/Taraxa-project/taraxa-indexer/internal/logging"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	migration "github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble/migrations"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	log "github.com/sirupsen/logrus"
)

var (
	http_port                        *int
	metrics_port                     *int
	blockchain_ws                    *string
	chain_id                         *int
	data_dir                         *string
	log_level                        *string
	yield_saving_interval            *int
	validators_yield_saving_interval *int
	sync_queue_limit                 *int
	signing_key                      *string
	oracle_address                   *string
	lara_address                     *string
)

func init() {
	http_port = flag.Int("http_port", 8080, "port to listen")
	metrics_port = flag.Int("metrics_port", 2112, "metrics http port")
	blockchain_ws = flag.String("blockchain_ws", "ws://localhost:8777", "ws url to connect to blockchain")
	chain_id = flag.Int("chain_id", 200, "chain id")
	data_dir = flag.String("data_dir", "./data", "path to directory where indexer database will be saved")
	log_level = flag.String("log_level", "info", "minimum log level. could be only [trace, debug, info, warn, error, fatal]")
	yield_saving_interval = flag.Int("yield_saving_interval", 100, "interval for saving total yield")
	validators_yield_saving_interval = flag.Int("validators_yield_saving_interval", 100, "interval for saving validators yield")
	sync_queue_limit = flag.Int("sync_queue_limit", 10, "limit of blocks in the sync queue")
	oracle_address = flag.String("oracle_address", "0x7EF7dB397007EdfBCFdefEE50Ff6B257D659E358", "oracles address")
	lara_address = flag.String("lara_address", "0xA188ECD2c5a4B0fC7Ce683bd69AeD1483f7e8fa0", "lara address")

	flag.Parse()

	logging.Config(filepath.Join(*data_dir, "logs"), *log_level)
	log.Print("\n\n\n")
	log.WithFields(log.Fields{
		"http_port":      *http_port,
		"blockchain_ws":  *blockchain_ws,
		"chain_id":       *chain_id,
		"signing_key":    *signing_key,
		"oracle_address": *oracle_address,
		"lara_address":   *lara_address,
		"data_dir":       *data_dir,
		"log_level":      *log_level}).
		Info("Application started")
}

func setupCloseHandler(st storage.Storage, fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go func() {
		<-c
		fn()
		os.Exit(0)
	}()
}

func main() {
	st := pebble.NewStorage(filepath.Join(*data_dir, "db"))
	setupCloseHandler(st, func() { st.Close() })

	swagger, err := api.GetSwagger()
	if err != nil {
		log.WithError(err).Fatal("Error loading swagger spec")
	}

	manager := migration.NewManager(st, *blockchain_ws)
	err = manager.ApplyAll()
	if err != nil {
		log.WithError(err).Fatal("Error applying migrations")
	}

	swagger.Servers = nil

	e := echo.New()

	e.Use(echomiddleware.OapiRequestValidator(swagger))
	// Add http error handler to return a proper error JSON on request error
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": err.Error()})
	}

	c := common.DefaultConfig()
	c.TotalYieldSavingInterval = uint64(*yield_saving_interval)
	c.ValidatorsYieldSavingInterval = uint64(*validators_yield_saving_interval)
	c.SyncQueueLimit = uint64(*sync_queue_limit)

	fin := st.GetFinalizationData()
	log.WithFields(log.Fields{"pbft_count": fin.PbftCount, "dag_count": fin.DagCount, "trx_count": fin.TrxCount}).Info("Loaded db with")

	apiHandler := api.NewApiHandler(st, c)
	api.RegisterHandlers(e, apiHandler)

	// Registers oracle cron
	if *signing_key == "" && *oracle_address == "" && *lara_address == "" {
		log.WithFields(log.Fields{"signing_key": *signing_key, "oracle_address": *oracle_address, "lara_address": *lara_address}).Fatal("Oracle address, Lara address and signing key should be both set but both empty")
	}

	indexer, wsClient := indexer.NewIndexer(*blockchain_ws, st, c)
	log.Info("Indexer initialized")
	rpc := ethclient.NewClient(wsClient.RpcClient())
	log.Info("RPC initialized")
	lara := lara.MakeLara(rpc, *signing_key, *lara_address, *oracle_address, *chain_id)
	log.Info("Lara initialized")
	o := oracle.MakeOracle(rpc, *signing_key, *oracle_address, *chain_id, *st)
	go oracle.RegisterCron(o, *yield_saving_interval)
	go lara.Run()
	go indexer.Run(*blockchain_ws, st, c, o)
	// start a http server for prometheus on a separate go routine
	go metrics.RunPrometheusServer(":" + strconv.FormatInt(int64(*metrics_port), 10))

	err = e.Start(":" + strconv.FormatInt(int64(*http_port), 10))
	log.WithError(err).Fatal("Can't start http server")
}
