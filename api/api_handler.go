//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml

package api

import (
	"net/http"
	"time"

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

func GetAddressDataPage[T storage.Paginated](a *ApiHandler, address AddressFilter, pag *PaginationParam) interface{} {
	logFields := log.Fields{"type": storage.GetTypeName[T](), "address": address, "pagination": pag}
	log.WithFields(logFields).Debug("GetAddressDataPage")

	ret, pagination := storage.GetObjectsPage[T](a.storage, address, getPaginationStart(pag.Start), pag.Limit)

	response := struct {
		PaginatedResponse
		Data []T `json:"data"`
	}{
		PaginatedResponse: *pagination,
		Data:              ret,
	}

	return response
}

// GetAddressDags returns all DAG blocks sent by the selected address
func (a *ApiHandler) GetAddressDags(ctx echo.Context, address AddressFilter, params GetAddressDagsParams) error {
	return ctx.JSON(http.StatusOK, GetAddressDataPage[Dag](a, address, &params.Pagination))
}

// GetAddressPbfts returns all PBFT blocks produced by the selected address
func (a *ApiHandler) GetAddressPbfts(ctx echo.Context, address AddressFilter, params GetAddressPbftsParams) error {
	return ctx.JSON(http.StatusOK, GetAddressDataPage[Pbft](a, address, &params.Pagination))
}

// GetAddressTransactions returns all transactions from and to the selected address
func (a *ApiHandler) GetAddressTransactions(ctx echo.Context, address AddressFilter, params GetAddressTransactionsParams) error {
	return ctx.JSON(http.StatusOK, GetAddressDataPage[Transaction](a, address, &params.Pagination))
}

// GetAddressPbftTotal returns total number of PBFT blocks produced for the selected address
func (a *ApiHandler) GetAddressStats(ctx echo.Context, address AddressFilter) error {
	log.WithField("address", address).Debug("GetAddressStats")

	addr := a.storage.GetAddressStats(address)
	return ctx.JSON(http.StatusOK, addr.StatsResponse)
}

// GetValidators returns all validators for the selected week and the number of PBFT blocks they produced
func (a *ApiHandler) GetValidators(ctx echo.Context, params GetValidatorsParams) error {
	log.WithField("params", params).Debug("GetValidators")
	year, week := getYearWeek(params.Week)
	stats := a.storage.GetWeekStats(year, week)
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
	log.WithField("params", params).Debug("GetValidatorsTotal")
	year, week := getYearWeek(params.Week)
	stats := a.storage.GetWeekStats(year, week)
	var count CountResponse
	count.Total = uint64(stats.Total)

	return ctx.JSON(http.StatusOK, count)
}

// GetValidator returns info about the validator for the selected week
func (a *ApiHandler) GetValidator(ctx echo.Context, address AddressParam, params GetValidatorParams) error {
	log.WithField("address", address).WithField("params", params).Debug("GetValidator")

	year, week := getYearWeek(params.Week)
	stats := a.storage.GetWeekStats(year, week)
	stats.Sort()

	validator := Validator{Address: address}

	for k, v := range stats.Validators {
		if v.Address == address {
			v.Rank = uint64(k + 1)
			validator = v
			break
		}
	}

	return ctx.JSON(http.StatusOK, validator)
}

func getPaginationStart(param *uint64) uint64 {
	if param == nil {
		return uint64(0)
	}

	return *param
}

func getYearWeek(week *WeekFilter) (int, int) {
	if week == nil {
		tn := time.Now()
		year, week := tn.ISOWeek()
		return year, week
	}

	return int(week.Year), int(week.Week)
}
