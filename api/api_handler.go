//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml

package api

import (
	"fmt"
	"net/http"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	. "github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/labstack/echo/v4"
)

type ApiHandler struct {
	Store *storage.Storage
}

func NewApiHandler(s *storage.Storage) *ApiHandler {
	return &ApiHandler{s}
}

// GetAddressDags returns all DAG blocks sent by the selected address
func (a *ApiHandler) GetAddressDags(ctx echo.Context, address AddressFilter, params GetAddressDagsParams) error {
	fmt.Println("GetAddressDags")
	var data []Dag
	response := struct {
		DagsPaginatedResponse
		Data []Dag
	}{
		Data: data,
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetAddressPbfts returns all PBFT blocks produced by the selected address
func (a *ApiHandler) GetAddressPbfts(ctx echo.Context, address AddressFilter, params GetAddressPbftsParams) error {
	fmt.Println("GetAddressPbfts")
	var data []Pbft
	response := struct {
		DagsPaginatedResponse
		Data []Pbft
	}{
		Data: data,
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetAddressTransactions returns all transactions from and to the selected address
func (a *ApiHandler) GetAddressTransactions(ctx echo.Context, address AddressFilter, params GetAddressTransactionsParams) error {
	fmt.Println("GetAddressTransactions")
	var data []Transaction
	response := struct {
		DagsPaginatedResponse
		Data []Transaction
	}{
		Data: data,
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetValidators returns all validators for the selected week and the number of PBFT blocks they produced
func (a *ApiHandler) GetValidators(ctx echo.Context, params GetValidatorsParams) error {
	fmt.Println("GetValidators")
	var data []Validator
	response := struct {
		DagsPaginatedResponse
		Data []Validator
	}{
		Data: data,
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetAddressDagTotal returns total number of DAG blocks sent from the selected address
func (a *ApiHandler) GetAddressDagTotal(ctx echo.Context, address AddressFilter) error {
	var addr storage.Address
	a.Store.GetFromDB(&addr, address)

	var count CountResponse
	count.Total = addr.DagTotal

	return ctx.JSON(http.StatusOK, count)
}

// GetAddressPbftTotal returns total number of PBFT blocks produced for the selected address
func (a *ApiHandler) GetAddressPbftTotal(ctx echo.Context, address AddressFilter) error {
	var addr storage.Address
	a.Store.GetFromDB(&addr, address)

	var count CountResponse
	count.Total = addr.PbftTotal

	return ctx.JSON(http.StatusOK, count)
}

// GetValidatorsTotal returns total number of PBFT blocks produced in selected week
func (a *ApiHandler) GetValidatorsTotal(ctx echo.Context, params GetValidatorsTotalParams) error {
	var count CountResponse
	count.Total = 0
	err := ctx.JSON(http.StatusOK, count)
	return err
}
