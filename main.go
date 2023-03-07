//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=models/models.cfg.yaml api/openapi.yaml

package main

import (
	"flag"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/api"
	"github.com/Taraxa-project/taraxa-indexer/internal/indexer"
	"github.com/Taraxa-project/taraxa-indexer/internal/logging"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var (
	http_port     *int
	blockchain_ws *string
	data_dir      *string
	log_level     *string
)

func init() {
	http_port = flag.Int("http_port", 8080, "port to listen")
	blockchain_ws = flag.String("blockchain_ws", "wss://ws.testnet.taraxa.io", "ws url to connect to blockchain")
	data_dir = flag.String("data_dir", "./data", "path to directory where indexer database will be saved")
	log_level = flag.String("log_level", "info", "minimum log level. could be only [trace, debug, info, warn, error, fatal]")

	flag.Parse()

	logging.Config(filepath.Join(*data_dir, "logs"), *log_level)
	log.Print("\n\n\n")
	log.WithFields(log.Fields{
		"http_port":     *http_port,
		"blockchain_ws": *blockchain_ws,
		"data_dir":      *data_dir,
		"log_level":     *log_level}).
		Info("Application started")
}

func setupCloseHandler(st *storage.Storage, fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go func() {
		<-c
		fn()
		os.Exit(0)
	}()
}

func main() {
	st := storage.NewStorage(filepath.Join(*data_dir, "db"))
	setupCloseHandler(st, func() { st.Close() })

	swagger, err := api.GetSwagger()
	if err != nil {
		log.WithError(err).Fatal("Error loading swagger spec")
	}

	swagger.Servers = nil

	e := echo.New()

	e.Use(middleware.OapiRequestValidator(swagger))
	// It is logging every incoming request. Do we need this?
	// e.Use(echomiddleware.Logger())

	apiHandler := api.NewApiHandler(st)
	api.RegisterHandlers(e, apiHandler)

	idx, err := indexer.NewIndexer(*blockchain_ws, st)
	if err != nil {
		log.WithError(err).Fatal("Can't create indexer")
	}
	go idx.Start()

	err = e.Start(":" + strconv.FormatInt(int64(*http_port), 10))
	log.WithError(err).Fatal("Can't start http server")
}
