//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml

package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
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

func GetHoldersDataPage(a *ApiHandler, pag *PaginationParam) interface{} {
	ret, pagination := storage.GetHoldersPage(a.storage, getPaginationStart(pag.Start), pag.Limit)

	response := struct {
		PaginatedResponse
		Data []models.Account `json:"data"`
	}{
		PaginatedResponse: *pagination,
		Data:              ret,
	}

	return response
}

func (a *ApiHandler) GetTransaction(ctx echo.Context, hash string) error {
	txHash := strings.ToLower(hash)

	tx := a.storage.GetTransactionByHash(txHash)
	if tx.Hash == "" {
		return ctx.JSON(http.StatusNotFound, "Transaction not found")
	}

	err := common.ProcessTransaction(&tx)
	if err != nil {
		log.WithError(err).WithField("hash", hash).Error("Error processing transaction")
	}

	return ctx.JSON(http.StatusOK, tx)
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

	// return last yield only for the last week stats
	last_year, last_week := getYearWeek(nil)
	if last_week == week && last_year == year {
		for i := 0; i < len(ret); i++ {
			resp, err := a.getAddressYield(ret[i].Address, nil)
			if err == nil {
				ret[i].Yield = resp.Yield
			}
		}
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

func (a *ApiHandler) GetHolders(ctx echo.Context, params GetHoldersParams) error {
	log.WithField("params", params).Debug("GetHolders")
	ret := GetHoldersDataPage(a, &params.Pagination)
	return ctx.JSON(http.StatusOK, ret)
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

	last_year, last_week := getYearWeek(nil)
	if last_week == week && last_year == year {
		resp, err := a.getAddressYield(address, nil)
		if err == nil {
			validator.Yield = resp.Yield
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

func (a *ApiHandler) getAddressYield(address AddressParam, block *uint64) (resp *models.YieldResponse, err error) {
	pbft_count := a.storage.GetFinalizationData().PbftCount
	block_num := common.GetYieldIntervalEnd(pbft_count, block, a.config.ValidatorsYieldSavingInterval)
	from_block := block_num - a.config.ValidatorsYieldSavingInterval + 1
	if pbft_count < block_num {
		err = fmt.Errorf("Not enough PBFT blocks(%d) to calculate yield for the interval [%d, %d]", pbft_count, from_block, block_num)
		return
	}
	return &models.YieldResponse{
		FromBlock: block_num - a.config.ValidatorsYieldSavingInterval + 1,
		ToBlock:   block_num,
		Yield:     a.storage.GetValidatorYield(address, block_num).Yield,
	}, nil
}

func (a *ApiHandler) GetAddressYield(ctx echo.Context, address AddressParam, params GetAddressYieldParams) error {
	resp, err := a.getAddressYield(address, params.BlockNumber)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (a *ApiHandler) GetTotalYield(ctx echo.Context, params GetTotalYieldParams) error {
	pbft_count := a.storage.GetFinalizationData().PbftCount
	block_num := common.GetYieldIntervalEnd(pbft_count, params.BlockNumber, a.config.TotalYieldSavingInterval)
	from_block := block_num - a.config.ValidatorsYieldSavingInterval + 1
	if pbft_count < block_num {
		return fmt.Errorf("Not enough PBFT blocks(%d) to calculate yield for the interval [%d, %d]", pbft_count, from_block, block_num)
	}
	resp := models.YieldResponse{
		FromBlock: block_num - a.config.TotalYieldSavingInterval + 1,
		ToBlock:   block_num,
		Yield:     a.storage.GetTotalYield(block_num).Yield,
	}
	return ctx.JSON(http.StatusOK, resp)
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
