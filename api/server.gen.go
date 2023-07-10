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
	// Returns yield for the address
	// (GET /address/{address}/yield)
	GetAddressYield(ctx echo.Context, address AddressParam, params GetAddressYieldParams) error
	// Returns the list of TARA token holders and their balances
	// (GET /holders)
	GetHolders(ctx echo.Context, params GetHoldersParams) error
	// Returns total supply
	// (GET /totalSupply)
	GetTotalSupply(ctx echo.Context) error
	// Returns total yield
	// (GET /totalYield)
	GetTotalYield(ctx echo.Context, params GetTotalYieldParams) error
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

// GetAddressYield converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressYield(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressYieldParams
	// ------------- Optional query parameter "blockNumber" -------------

	err = runtime.BindQueryParameter("form", true, false, "blockNumber", ctx.QueryParams(), &params.BlockNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter blockNumber: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAddressYield(ctx, address, params)
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

// GetTotalYield converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalYield(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTotalYieldParams
	// ------------- Optional query parameter "blockNumber" -------------

	err = runtime.BindQueryParameter("form", true, false, "blockNumber", ctx.QueryParams(), &params.BlockNumber)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter blockNumber: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTotalYield(ctx, params)
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
	router.GET(baseURL+"/address/:address/yield", wrapper.GetAddressYield)
	router.GET(baseURL+"/holders", wrapper.GetHolders)
	router.GET(baseURL+"/totalSupply", wrapper.GetTotalSupply)
	router.GET(baseURL+"/totalYield", wrapper.GetTotalYield)
	router.GET(baseURL+"/transaction/:hash/internal_transactions", wrapper.GetInternalTransactions)
	router.GET(baseURL+"/transaction/:hash/logs", wrapper.GetTransactionLogs)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)
	router.GET(baseURL+"/validators/:address", wrapper.GetValidator)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"S+saVn1IRWaZS+lu+78AAAD//yhZvYg/OwAA",
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
