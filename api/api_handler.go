//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml

package api

import (
	"net/http"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	. "github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/labstack/echo/v4"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type ApiHandler struct {
	storage storage.Storage
	config  *common.Config
}

func NewApiHandler(s storage.Storage, c *common.Config) *ApiHandler {
	return &ApiHandler{s, c}
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

	date, _ := goment.New()
	date.SetISOWeek(int(week))
	date.SetISOWeekYear(int(year))

	tn, _ := goment.New()

	w := WeekResponse{
		Week:      &week,
		Year:      &year,
		StartDate: uint64(date.StartOf("week").ToTime().Unix()),
		EndDate:   uint64(date.EndOf("week").ToTime().Unix()),
		HasNext:   tn.ISOWeekYear() != int(year) || tn.ISOWeek() != int(week),
	}

	response := struct {
		ValidatorsPaginatedResponse
		Data []Validator  `json:"data"`
		Week WeekResponse `json:"week"`
	}{
		ValidatorsPaginatedResponse: *pagination,
		Data:                        ret,
		Week:                        w,
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
	address = strings.ToLower(address)
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

func (a *ApiHandler) GetTotalSupply(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, a.storage.GetTotalSupply().String())
}

func (a *ApiHandler) GetInternalTransactions(ctx echo.Context, hash HashParam) error {
	return ctx.JSON(http.StatusOK, a.storage.GetInternalTransactions(hash))
}

func (a *ApiHandler) GetTransactionLogs(ctx echo.Context, hash HashParam) error {
	return ctx.JSON(http.StatusOK, a.storage.GetTransactionLogs(hash))
}

func (a *ApiHandler) GetAddressYield(ctx echo.Context, address AddressParam, params GetAddressYieldParams) error {
	block_num := common.GetYieldIntervalEnd(a.storage, params.BlockNumber, a.config.ValidatorsYieldSavingInterval)
	return ctx.JSON(http.StatusOK, a.storage.GetValidatorYield(address, block_num))
}

func (a *ApiHandler) GetTotalYield(ctx echo.Context, params GetTotalYieldParams) error {
	block_num := common.GetYieldIntervalEnd(a.storage, params.BlockNumber, a.config.TotalYieldSavingInterval)
	return ctx.JSON(http.StatusOK, a.storage.GetTotalYield(block_num))
}

func getPaginationStart(param *uint64) uint64 {
	if param == nil {
		return uint64(0)
	}

	return *param
}

func getYearWeek(w *WeekParam) (int32, int32) {
	if w == nil || w.Week == nil || w.Year == nil {
		tn, _ := goment.New()
		return int32(tn.ISOWeekYear()), int32(tn.ISOWeek())
	}

	return int32(*w.Year), int32(*w.Week)
}
