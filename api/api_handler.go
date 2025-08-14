//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml openapi.yaml

package api

import (
	"bytes"
	"strconv"
	"time"

	"fmt"
	"net/http"
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	. "github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/labstack/echo/v4"
	"github.com/nleeper/goment"
	log "github.com/sirupsen/logrus"
)

type ApiHandler struct {
	storage storage.Storage
	config  *common.Config
	stats   *chain.Stats
}

func NewApiHandler(s storage.Storage, c *common.Config, stats *chain.Stats) *ApiHandler {
	return &ApiHandler{s, c, stats}
}

func formatAddress(address AddressParam) (formatted string, err error) {
	// starts with 0x
	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}
	if len(address) != 42 {
		err = fmt.Errorf("invalid address length")
	}
	formatted = strings.ToLower(address)
	return
}

func GetAddressDataPage[T storage.Paginated](a *ApiHandler, address AddressFilter, pag *PaginationParam) any {
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

func GetHoldersDataPage(a *ApiHandler, pag *PaginationParam) any {
	ret, pagination := storage.GetHoldersPage(a.storage, getPaginationStart(pag.Start), pag.Limit)

	response := struct {
		PaginatedResponse
		Data []Account `json:"data"`
	}{
		PaginatedResponse: *pagination,
		Data:              ret,
	}

	return response
}

func (a *ApiHandler) GetChainStats(ctx echo.Context) error {
	stats := a.stats.GetStats()
	if stats == nil {
		return ctx.JSON(http.StatusNotFound, "Chain stats not found")
	}
	return ctx.JSON(http.StatusOK, *stats)
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
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetAddressDataPage[Dag](a, address, &params.Pagination))
}

// GetAddressPbfts returns all PBFT blocks produced by the selected address
func (a *ApiHandler) GetAddressPbfts(ctx echo.Context, address AddressFilter, params GetAddressPbftsParams) error {
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetAddressDataPage[Pbft](a, address, &params.Pagination))
}

// GetAddressTransactions returns all transactions from and to the selected address
func (a *ApiHandler) GetAddressTransactions(ctx echo.Context, address AddressFilter, params GetAddressTransactionsParams) error {
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetAddressDataPage[Transaction](a, address, &params.Pagination))
}

// GetAddressPbftTotal returns total number of PBFT blocks produced for the selected address
func (a *ApiHandler) GetAddressStats(ctx echo.Context, address AddressFilter) error {
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
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
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
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

	statsResponse := a.storage.GetAddressStats(address).StatsResponse
	validator.RegistrationBlock = statsResponse.ValidatorRegisteredBlock

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

func (a *ApiHandler) getAddressYield(address AddressParam, block *uint64) (resp *YieldResponse, err error) {
	pbft_count := a.storage.GetFinalizationData().PbftCount
	block_num := common.GetYieldIntervalEnd(pbft_count, block, a.config.ValidatorsYieldSavingInterval)
	from_block := block_num - a.config.ValidatorsYieldSavingInterval + 1
	if pbft_count < block_num {
		err = fmt.Errorf("not enough PBFT blocks(%d) to calculate yield for the interval [%d, %d]", pbft_count, from_block, block_num)
		return
	}
	return &YieldResponse{
		FromBlock: block_num - a.config.ValidatorsYieldSavingInterval + 1,
		ToBlock:   block_num,
		Yield:     a.storage.GetValidatorYield(address, block_num).Yield,
	}, nil
}

func (a *ApiHandler) GetAddressYield(ctx echo.Context, address AddressParam, params GetAddressYieldParams) error {
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	resp, err := a.getAddressYield(address, params.BlockNumber)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, resp)
}
func (a *ApiHandler) GetAddressYieldForInterval(ctx echo.Context, address AddressParam, params GetAddressYieldForIntervalParams) error {
	pbft_count := a.storage.GetFinalizationData().PbftCount
	block_num := common.GetYieldIntervalEnd(pbft_count, params.FromBlock, a.config.ValidatorsYieldSavingInterval)

	from_block := block_num - a.config.ValidatorsYieldSavingInterval + 1
	to_block := common.GetYieldIntervalEnd(pbft_count, &params.ToBlock, a.config.ValidatorsYieldSavingInterval)

	prefix := []byte(pebble.GetPrefixKey(pebble.GetPrefix(storage.Yield{}), address))
	yield := float64(0)
	count := int64(0)
	a.storage.ForEachFromKey(prefix, []byte(storage.FormatIntToKey(from_block)), storage.Forward, func(key, res []byte) (stop bool) {
		if bytes.Equal(key, bytes.Join([][]byte{prefix, []byte(storage.FormatIntToKey(to_block))}, []byte(""))) {
			return true
		}
		y := storage.Yield{}
		err := rlp.DecodeBytes(res, &y)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return false
		}

		fmt.Println("Key: ", string(key))
		fmt.Printf("Yield: %f, adding: %f\n", yield, common.ParseFloat(y.Yield))

		yield += common.ParseFloat(y.Yield)
		count++

		return false
	})
	if count == 0 {
		return fmt.Errorf("no yield data found for the %s at interval [%d, %d]", address, from_block, to_block)
	}
	yield /= float64(count)

	resp := YieldResponse{
		FromBlock: from_block,
		ToBlock:   to_block,
		Yield:     common.FormatFloat(yield),
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (a *ApiHandler) GetTotalYield(ctx echo.Context, params GetTotalYieldParams) error {
	pbft_count := a.storage.GetFinalizationData().PbftCount
	block_num := common.GetYieldIntervalEnd(pbft_count, params.BlockNumber, a.config.TotalYieldSavingInterval)
	from_block := block_num - a.config.ValidatorsYieldSavingInterval + 1
	if pbft_count < block_num {
		return fmt.Errorf("not enough PBFT blocks(%d) to calculate yield for the interval [%d, %d]", pbft_count, from_block, block_num)
	}
	resp := YieldResponse{
		FromBlock: block_num - a.config.TotalYieldSavingInterval + 1,
		ToBlock:   block_num,
		Yield:     a.storage.GetTotalYield(block_num).Yield,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (a *ApiHandler) GetLatestMonthlyActiveAddresses() (resp MonthlyActiveAddressesResponse) {
	ma := storage.MonthlyActiveAddresses{}
	a.storage.ForEach(&ma, "", nil, storage.Backward, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &ma)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return
		}

		parts := strings.Split(string(res), "|")

		to_date, err := strconv.ParseUint(parts[1], 10, 64)
		if err != nil {
			log.WithField("to_date_raw", parts[1]).Fatal("LatestMonthlyActiveAddresses: error parsing uint")
		}

		resp.Count = ma.Count
		resp.FromDate = to_date
		resp.ToDate = to_date - common.Days30

		return true
	})
	return
}

func (a *ApiHandler) GetMonthlyActiveAddresses(ctx echo.Context, params GetMonthlyActiveAddressesParams) error {
	from_date, to_date := common.MonthInterval(params.Date)

	log.WithField("from_date", from_date).WithField("to_date", to_date).Debug("GetMonthlyActiveAddresses")

	count, err := storage.GetMonthlyActiveAddresses(a.storage, from_date, to_date)
	if err != nil {
		if params.Date == nil {
			return ctx.JSON(http.StatusOK, a.GetLatestMonthlyActiveAddresses())
		}
		return ctx.JSON(http.StatusRequestTimeout, err.Error())
	}

	resp := MonthlyActiveAddressesResponse{
		Count:    count,
		FromDate: from_date,
		ToDate:   to_date,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (a *ApiHandler) GetMonthlyStats(ctx echo.Context, params GetMonthlyStatsParams) error {
	from_date, to_date := common.MonthInterval(params.Date)

	totalStats := storage.EmptyTrxGasStats()
	stats := storage.EmptyTrxGasStats()
	count := int64(0)

	a.storage.ForEach(&stats, "", nil, storage.Forward, func(key []byte, res []byte) (stop bool) {
		ts := storage.GetTimestampFromKey(key)
		if ts < from_date || ts > to_date {
			return false
		}
		err := rlp.DecodeBytes(res, &stats)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return false
		}

		totalStats.Add(stats)
		count++

		return count == 30
	})

	if count < 30 {
		return ctx.JSON(http.StatusNotFound, "Not enough stats found for the interval")
	}

	return ctx.JSON(http.StatusOK, MonthlyStatsResponse{
		FromDate: from_date,
		ToDate:   to_date,
		GasUsed:  totalStats.GasUsed.String(),
		TrxCount: totalStats.TrxCount,
	})
}

func (a *ApiHandler) GetContractStats(ctx echo.Context, params GetContractStatsParams) error {
	contracts := []ContractStatsResponse{}
	stats := storage.AddressStats{}
	start := time.Now()
	a.storage.ForEach(&stats, "", nil, storage.Forward, func(key []byte, res []byte) (stop bool) {
		err := rlp.DecodeBytes(res, &stats)
		if err != nil {
			log.WithError(err).Fatal("Error decoding data from db")
			return false
		}

		if stats.ContractRegisteredTimestamp == nil {
			return false
		}

		count := storage.ReceivedTransactionsCount(a.storage, stats.Address, params.FromDate, params.ToDate)

		if count == 0 {
			return false
		}

		contracts = append(contracts, ContractStatsResponse{
			Address:           stats.Address,
			CreationDate:      *stats.ContractRegisteredTimestamp,
			TransactionsCount: count,
		})
		return false
	})
	log.WithField("time", time.Since(start)).Info("GetContractStats")
	return ctx.JSON(http.StatusOK, contracts)
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

func (a *ApiHandler) GetMonthlyAverageDailyActiveWallets(ctx echo.Context, address AddressParam, params GetMonthlyAverageDailyActiveWalletsParams) error {
	address, err := formatAddress(address)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	lastDayStart := common.DayStart(uint64(time.Now().Unix()))
	if params.Date != nil {
		lastDayStart = common.DayStart(*params.Date)
	}
	average := a.calculateContract30DayAverage(address, lastDayStart)

	return ctx.JSON(http.StatusOK, average)
}

func (a *ApiHandler) calculateContract30DayAverage(contractAddress string, lastDayStart uint64) MonthlyAverageDailyActiveWalletsResponse {
	startDay := lastDayStart - common.Days30

	var totalUsers uint64 = 0
	const totalDays = 30

	// Iterate through the last 30 days
	for day := startDay; day < lastDayStart; day += common.Day {
		users := a.storage.GetDailyContractUsers(contractAddress, day)
		// Count users for this day (0 if no data)
		totalUsers += uint64(len(users.Users))
	}

	// Calculate average including days with no data (divide by 30)
	var average float32 = float32(totalUsers) / float32(totalDays)

	return MonthlyAverageDailyActiveWalletsResponse{
		Count: average,
	}
}
