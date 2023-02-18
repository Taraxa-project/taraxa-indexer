package main

import (
	"fmt"
	"os"

	"github.com/Taraxa-project/taraxa-indexer/api"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	apiHandler := api.NewApiHandler()
	api.RegisterHandlers(e, apiHandler)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
