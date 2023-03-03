//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml

package api

import (
	"net/http"

	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	. "github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ApiHandler struct {
	storage *storage.Storage
}

func NewApiHandler(s *storage.Storage) *ApiHandler {
	return &ApiHandler{s}
}

// GetAddressDags returns all DAG blocks sent by the selected address
func (a *ApiHandler) GetAddressDags(ctx echo.Context, address AddressFilter, params GetAddressDagsParams) error {
	logFields := log.Fields{"address": address, "params": params}
	log.WithFields(logFields).Trace("GetAddressDags")

	ret, pagination, err := storage.GetObjectsPage[Dag](a.storage, address, getPaginationStart(params.Pagination.Start), params.Pagination.Limit)
	if err != nil {
		logFields["error"] = err
		log.WithFields(logFields).Fatal("Error getting Dags")
	}

	stats := a.storage.GetAddressStats(address)
	pagination.Total = stats.DagsCount

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
	logFields := log.Fields{"address": address, "params": params}
	log.WithFields(logFields).Trace("GetAddressPbfts")

	ret, pagination, err := storage.GetObjectsPage[Pbft](a.storage, address, getPaginationStart(params.Pagination.Start), params.Pagination.Limit)
	if err != nil {
		logFields["error"] = err
		log.WithFields(logFields).Fatal("Error getting Pbfts")
	}

	stats := a.storage.GetAddressStats(address)
	pagination.Total = stats.PbftCount

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
	logFields := log.Fields{"address": address, "params": params}
	log.WithFields(logFields).Trace("GetAddressTransactions")

	ret, pagination, err := storage.GetObjectsPage[Transaction](a.storage, address, getPaginationStart(params.Pagination.Start), params.Pagination.Limit)
	if err != nil {
		logFields["error"] = err
		log.WithFields(logFields).Debug("Error getting Transactions")
	}

	stats := a.storage.GetAddressStats(address)
	pagination.Total = stats.TransactionsCount

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
	log.WithField("address", address).Trace("GetAddressStats")

	addr := a.storage.GetAddressStats(address)
	return ctx.JSON(http.StatusOK, addr.StatsResponse)
}

// GetValidators returns all validators for the selected week and the number of PBFT blocks they produced
func (a *ApiHandler) GetValidators(ctx echo.Context, params GetValidatorsParams) error {
	log.WithField("params", params).Trace("GetValidators")

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
	log.WithField("params", params).Trace("GetValidatorsTotal")
	stats := a.storage.GetWeekStats(int(params.Filter.Year), int(params.Filter.Week))
	var count CountResponse
	count.Total = uint64(stats.Total)

	return ctx.JSON(http.StatusOK, count)
}

func getPaginationStart(param *uint64) uint64 {
	if param == nil {
		return uint64(0)
	}

	return *param
}
