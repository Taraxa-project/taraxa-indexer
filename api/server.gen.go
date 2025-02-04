// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
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
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
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
	// Returns yield for the address
	// (GET /address/{address}/yieldForInterval)
	GetAddressYieldForInterval(ctx echo.Context, address AddressParam, params GetAddressYieldForIntervalParams) error
	// Returns chain stats
	// (GET /chainStats)
	GetChainStats(ctx echo.Context) error
	// Returns the list of TARA token holders and their balances
	// (GET /holders)
	GetHolders(ctx echo.Context, params GetHoldersParams) error
	// Returns total supply
	// (GET /totalSupply)
	GetTotalSupply(ctx echo.Context) error
	// Returns total yield
	// (GET /totalYield)
	GetTotalYield(ctx echo.Context, params GetTotalYieldParams) error
	// Returns the decoded transaction
	// (GET /transaction/{hash})
	GetTransaction(ctx echo.Context, hash HashParam) error
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

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressDags(ctx, address, params)
	return err
}

// GetAddressPbfts converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressPbfts(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressPbfts(ctx, address, params)
	return err
}

// GetAddressStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressStats(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressStats(ctx, address)
	return err
}

// GetAddressTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressTransactions(ctx, address, params)
	return err
}

// GetAddressYield converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressYield(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressYield(ctx, address, params)
	return err
}

// GetAddressYieldForInterval converts echo context to params.
func (w *ServerInterfaceWrapper) GetAddressYieldForInterval(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetAddressYieldForIntervalParams
	// ------------- Optional query parameter "fromBlock" -------------

	err = runtime.BindQueryParameter("form", true, false, "fromBlock", ctx.QueryParams(), &params.FromBlock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fromBlock: %s", err))
	}

	// ------------- Required query parameter "toBlock" -------------

	err = runtime.BindQueryParameter("form", true, true, "toBlock", ctx.QueryParams(), &params.ToBlock)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter toBlock: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetAddressYieldForInterval(ctx, address, params)
	return err
}

// GetChainStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetChainStats(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetChainStats(ctx)
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetHolders(ctx, params)
	return err
}

// GetTotalSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetTotalSupply(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTotalYield(ctx, params)
	return err
}

// GetTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransaction(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithOptions("simple", "hash", ctx.Param("hash"), &hash, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTransaction(ctx, hash)
	return err
}

// GetInternalTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetInternalTransactions(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithOptions("simple", "hash", ctx.Param("hash"), &hash, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetInternalTransactions(ctx, hash)
	return err
}

// GetTransactionLogs converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransactionLogs(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "hash" -------------
	var hash HashParam

	err = runtime.BindStyledParameterWithOptions("simple", "hash", ctx.Param("hash"), &hash, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter hash: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
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

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetValidatorsTotal(ctx, params)
	return err
}

// GetValidator converts echo context to params.
func (w *ServerInterfaceWrapper) GetValidator(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address AddressParam

	err = runtime.BindStyledParameterWithOptions("simple", "address", ctx.Param("address"), &address, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
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

	// Invoke the callback with all the unmarshaled arguments
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
	router.GET(baseURL+"/address/:address/yieldForInterval", wrapper.GetAddressYieldForInterval)
	router.GET(baseURL+"/chainStats", wrapper.GetChainStats)
	router.GET(baseURL+"/holders", wrapper.GetHolders)
	router.GET(baseURL+"/totalSupply", wrapper.GetTotalSupply)
	router.GET(baseURL+"/totalYield", wrapper.GetTotalYield)
	router.GET(baseURL+"/transaction/:hash", wrapper.GetTransaction)
	router.GET(baseURL+"/transaction/:hash/internal_transactions", wrapper.GetInternalTransactions)
	router.GET(baseURL+"/transaction/:hash/logs", wrapper.GetTransactionLogs)
	router.GET(baseURL+"/validators", wrapper.GetValidators)
	router.GET(baseURL+"/validators/total", wrapper.GetValidatorsTotal)
	router.GET(baseURL+"/validators/:address", wrapper.GetValidator)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"hZRuNv8LAAD//7na06oqRQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
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
	res := make(map[string]func() ([]byte, error))
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
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
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
