//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=models/models.cfg.yaml api/openapi.yaml

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/api"
	"github.com/Taraxa-project/taraxa-indexer/internal/indexer"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

var (
	http_port     *int
	blockchain_ws *string
	db_path       *string
)

func init() {
	http_port = flag.Int("http_port", 8080, "port to listen")
	blockchain_ws = flag.String("blockchain_ws", "wss://ws.testnet.taraxa.io", "ws url to connect to blockchain")
	db_path = flag.String("db_path", "./data", "path to directory where indexer database will be saved")
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

func GetCommit() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}

func main() {
	flag.Parse()
	fmt.Println("Built from commit", GetCommit())
	fmt.Println("passed blockchain_ws", *blockchain_ws)
	fmt.Println("passed db_path", *db_path)

	st := storage.NewStorage(*db_path)

	setupCloseHandler(st, func() { st.Close() })

	e := echo.New()

	e.Use(echomiddleware.Logger())

	apiHandler := api.NewApiHandler(st)
	api.RegisterHandlers(e, apiHandler)

	idx, err := indexer.NewIndexer(*blockchain_ws, st)
	if err != nil {
		log.Fatal("Problem with indexer", err)
	}
	go idx.Start()
	e.Logger.Fatal(e.Start(":" + strconv.FormatInt(int64(*http_port), 10)))
}
