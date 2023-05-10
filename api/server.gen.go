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
	// Returns total supply
	// (GET /totalSupply)
	GetTotalSupply(ctx echo.Context) error
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

// GetTotalSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalSupply(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTotalSupply(ctx)
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
	router.GET(baseURL+"/totalSupply", wrapper.GetTotalSupply)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)
	router.GET(baseURL+"/validators/:address", wrapper.GetValidator)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"t0dDKQAA",
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
