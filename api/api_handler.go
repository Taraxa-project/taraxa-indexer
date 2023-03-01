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
	storage *storage.Storage
}

func NewApiHandler(s *storage.Storage) *ApiHandler {
	return &ApiHandler{s}
}

// GetAddressDags returns all DAG blocks sent by the selected address
func (a *ApiHandler) GetAddressDags(ctx echo.Context, address AddressFilter, params GetAddressDagsParams) error {
	fmt.Println("GetAddressDags")
	stats, err := a.storage.GetAddressStats(address)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data for address "+address+" not found")
	}

	ret, pagination, err := storage.GetObjectsPage[Dag](a.storage, address, getPaginationStart(params.Pagination.Start), params.Pagination.Limit)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Dags not found for "+address)
	}

	pagination.Total = stats.DagTotal
	response := struct {
		PaginatedResponse
		Data []Dag `json:"data"`
	}{
		PaginatedResponse: *pagination,
		Data:              ret,
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetAddressPbfts returns all PBFT blocks produced by the selected address
func (a *ApiHandler) GetAddressPbfts(ctx echo.Context, address AddressFilter, params GetAddressPbftsParams) error {
	fmt.Println("GetAddressPbfts")
	stats, err := a.storage.GetAddressStats(address)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data for address "+address+" not found")
	}

	ret, pagination, err := storage.GetObjectsPage[Pbft](a.storage, address, getPaginationStart(params.Pagination.Start), params.Pagination.Limit)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Pbfts not found for "+address)
	}

	pagination.Total = stats.PbftTotal
	response := struct {
		PaginatedResponse
		Data []Pbft `json:"data"`
	}{
		PaginatedResponse: *pagination,
		Data:              ret,
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetAddressTransactions returns all transactions from and to the selected address
func (a *ApiHandler) GetAddressTransactions(ctx echo.Context, address AddressFilter, params GetAddressTransactionsParams) error {
	fmt.Println("GetAddressTransactions")
	stats, err := a.storage.GetAddressStats(address)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Data for address "+address+" not found")
	}

	ret, pagination, err := storage.GetObjectsPage[Transaction](a.storage, address, getPaginationStart(params.Pagination.Start), params.Pagination.Limit)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "Transactions not found for "+address)
	}

	pagination.Total = stats.TxTotal
	response := struct {
		PaginatedResponse
		Data []Transaction `json:"data"`
	}{
		PaginatedResponse: *pagination,
		Data:              ret,
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetAddressPbftTotal returns total number of PBFT blocks produced for the selected address
func (a *ApiHandler) GetAddressStats(ctx echo.Context, address AddressFilter) error {
	addr, err := a.storage.GetAddressStats(address)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusNotFound, err)
	}
	var count StatsResponse
	count.PbftCount = addr.PbftTotal
	count.DagsCount = addr.DagTotal
	count.TransactionsCount = addr.TxTotal

	return ctx.JSON(http.StatusOK, count)
}

// GetValidators returns all validators for the selected week and the number of PBFT blocks they produced
func (a *ApiHandler) GetValidators(ctx echo.Context, params GetValidatorsParams) error {
	fmt.Println("GetValidators")

	stats := a.storage.GetWeekStats(int(params.Week.Year), int(params.Week.Week))
	ret, pagination := stats.GetPaginated(getPaginationStart(params.Pagination.Start), params.Pagination.Limit)

	response := struct {
		DagsPaginatedResponse
		Data []Validator `json:"data"`
	}{
		DagsPaginatedResponse: *pagination,
		Data:                  ret,
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetValidatorsTotal returns total number of PBFT blocks produced in selected week
func (a *ApiHandler) GetValidatorsTotal(ctx echo.Context, params GetValidatorsTotalParams) error {
	stats := a.storage.GetWeekStats(int(params.Filter.Year), int(params.Filter.Week))
	var count CountResponse
	count.Total = uint64(stats.Total)
	err := ctx.JSON(http.StatusOK, count)
	return err
}

func getPaginationStart(param *uint64) uint64 {
	if param == nil {
		return uint64(0)
	}

	return *param
}
