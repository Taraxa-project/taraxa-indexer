// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.5-0.20230118012357-f4cf8f9a5703 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all DAG blocks
	// (GET /address/{address}/dags)
	GetAddressDags(ctx echo.Context, address AddressParam, params GetAddressDagsParams) error
	// Returns all PBFT blocks
	// (GET /address/{address}/pbfts)
	GetAddressPbfts(ctx echo.Context, address AddressParam, params GetAddressPbftsParams) error
	// Returns stats for the address
	// (GET /address/{address}/stats)
	GetAddressStats(ctx echo.Context, address AddressParam) error
	// Returns all transactions
	// (GET /address/{address}/transactions)
	GetAddressTransactions(ctx echo.Context, address AddressParam, params GetAddressTransactionsParams) error
	// Returns all validators
	// (GET /validators)
	GetValidators(ctx echo.Context, params GetValidatorsParams) error
	// Returns total number of PBFT blocks
	// (GET /validators/total)
	GetValidatorsTotal(ctx echo.Context, params GetValidatorsTotalParams) error
	// Returns info about the validator
	// (GET /validators/{address})
	GetValidator(ctx echo.Context, address AddressParam, params GetValidatorParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAddressDags converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressDags(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressDagsParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressDags(ctx, address, params)
	return err
}

// GetAddressPbfts converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressPbfts(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressPbftsParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressPbfts(ctx, address, params)
	return err
}

// GetAddressStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressStats(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressStats(ctx, address)
	return err
}

// GetAddressTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressTransactionsParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressTransactions(ctx, address, params)
	return err
}

// GetValidators converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidators(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorsParams
	// ------------- Optional query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, false, "week", ctx.QueryParams(), &params.Week)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter week: %s", err))
	}

	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidators(ctx, params)
	return err
}

// GetValidatorsTotal converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidatorsTotal(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorsTotalParams
	// ------------- Optional query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, false, "week", ctx.QueryParams(), &params.Week)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter week: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidatorsTotal(ctx, params)
	return err
}

// GetValidator converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidator(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetValidatorParams
	// ------------- Optional query parameter "week" -------------

	err = runtime.BindQueryParameter("form", true, false, "week", ctx.QueryParams(), &params.Week)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter week: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetValidator(ctx, address, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/address/:address/dags", wrapper.GetAddressDags)
	router.GET(baseURL+"/address/:address/pbfts", wrapper.GetAddressPbfts)
	router.GET(baseURL+"/address/:address/stats", wrapper.GetAddressStats)
	router.GET(baseURL+"/address/:address/transactions", wrapper.GetAddressTransactions)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)
	router.GET(baseURL+"/validators/:address", wrapper.GetValidator)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"3UdUaJe5kh53/wsAAP//Vyg+omMnAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
