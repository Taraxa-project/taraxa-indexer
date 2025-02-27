// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
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
	"AA==",
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
