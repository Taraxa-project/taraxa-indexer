//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=models/models.cfg.yaml api/openapi.yaml

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/api"
	"github.com/Taraxa-project/taraxa-indexer/internal/indexer"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var (
	http_port     *int
	blockchain_ws *string
)

func init() {
	http_port = flag.Int("http_port", 8080, "port to listen")
	blockchain_ws = flag.String("blockchain_ws", "wss://ws.testnet.taraxa.io", "ws url to connect to blockchain")
}

func setupCloseHandler(st *storage.Storage, fn func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go func() {
		<-c
		fn()
		os.Exit(0)
	}()
}

func main() {
	st := storage.NewStorage("./data/indexer.db")

	setupCloseHandler(st, func() { st.Close() })

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	apiHandler := api.NewApiHandler(st)
	api.RegisterHandlers(e, apiHandler)
	flag.Parse()
	fmt.Println("passed blockchain_ws", *blockchain_ws)

	st.RecordFinalizedPeriod(65000)
	idx, err := indexer.NewIndexer(*blockchain_ws, st)
	if err != nil {
		log.Fatal("Problem with indexer", err)
	}
	go idx.Start()
	e.Logger.Fatal(e.Start(":" + strconv.FormatInt(int64(*http_port), 10)))
}
