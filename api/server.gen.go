// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
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
	// Returns the list of TARA token holders and their balances
	// (GET /holders)
	GetHolders(ctx echo.Context, params GetHoldersParams) error
	// Returns total supply
	// (GET /totalSupply)
	GetTotalSupply(ctx echo.Context) error
	// Returns internal transactions
	// (GET /transaction/{hash}/internal_transactions)
	GetInternalTransactions(ctx echo.Context, hash HashParam) error
	// Returns event logs of transaction
	// (GET /transaction/{hash}/logs)
	GetTransactionLogs(ctx echo.Context, hash HashParam) error
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

// GetHolders converts echo context to params.
func (w *ServerInterfaceWrapper) GetHolders(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetHoldersParams
	// ------------- Required query parameter "pagination" -------------

	err = runtime.BindQueryParameter("form", true, true, "pagination", ctx.QueryParams(), &params.Pagination)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter pagination: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetHolders(ctx, params)
	return err
}

// GetTotalSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalSupply(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTotalSupply(ctx)
	return err
}

// GetInternalTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetInternalTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "hash", runtime.ParamLocationPath, ctx.Param("hash"), &hash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetInternalTransactions(ctx, hash)
	return err
}

// GetTransactionLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransactionLogs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "hash", runtime.ParamLocationPath, ctx.Param("hash"), &hash)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTransactionLogs(ctx, hash)
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
	router.GET(baseURL+"/holders", wrapper.GetHolders)
	router.GET(baseURL+"/totalSupply", wrapper.GetTotalSupply)
	router.GET(baseURL+"/transaction/:hash/internal_transactions", wrapper.GetInternalTransactions)
	router.GET(baseURL+"/transaction/:hash/logs", wrapper.GetTransactionLogs)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)
	router.GET(baseURL+"/validators/:address", wrapper.GetValidator)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
