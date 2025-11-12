package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage"
	"github.com/Taraxa-project/taraxa-indexer/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Simple test that focuses on the handler logic without complex mocking
func TestGetContractsHandler_Direct(t *testing.T) {
	// Create a simple mock storage that implements minimal interface
	mockStorage := &SimpleStorage{}

	// Create API handler directly
	config := common.DefaultConfig()
	apiHandler := &ApiHandler{
		storage: mockStorage,
		config:  config,
		stats:   nil,
	}

	// Test the handler directly
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := apiHandler.GetContracts(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse response
	var response models.ContractsResponse
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Should return empty array, not null
	assert.NotNil(t, response.Contracts)
	assert.Equal(t, 0, len(response.Contracts))

	// Verify JSON contains empty array
	assert.Contains(t, rec.Body.String(), `"contracts":[]`)
}

func TestGetContractStatsHandler_Direct(t *testing.T) {
	// Create a simple mock storage
	mockStorage := &SimpleStorage{}

	// Create API handler directly
	config := common.DefaultConfig()
	apiHandler := &ApiHandler{
		storage: mockStorage,
		config:  config,
		stats:   nil,
	}

	// Test the handler directly
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/contractStats?fromDate=1704067200&toDate=1735689600", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Manually set the parsed parameters (normally done by generated server code)
	params := models.GetContractStatsParams{
		FromDate: 1704067200,
		ToDate:   1735689600,
	}

	err := apiHandler.GetContractStats(c, params)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Should return empty array
	assert.Contains(t, rec.Body.String(), "[]")
}

// SimpleStorage is a minimal storage implementation for testing
type SimpleStorage struct{}

func (s *SimpleStorage) Clean() error { return nil }
func (s *SimpleStorage) Close() error { return nil }

func (s *SimpleStorage) ForEach(o any, key_prefix string, start *uint64, direction storage.Direction, fn func(key, res []byte) (stop bool)) {
	// For testing, don't call the callback - simulates empty result set
}

func (s *SimpleStorage) ForEachFromKey(prefix, start_key []byte, direction storage.Direction, fn func(key, res []byte) (stop bool)) {
}
func (s *SimpleStorage) NewBatch() storage.Batch                         { return nil }
func (s *SimpleStorage) GetTotalSupply() *storage.TotalSupply            { return nil }
func (s *SimpleStorage) GetAccounts() storage.Accounts                   { return storage.Accounts{} }
func (s *SimpleStorage) GetWeekStats(year, week int32) storage.WeekStats { return storage.WeekStats{} }
func (s *SimpleStorage) GetFinalizationData() *common.FinalizationData {
	return &common.FinalizationData{}
}
func (s *SimpleStorage) GetDayStats(timestamp uint64) storage.DayStatsWithTimestamp {
	return storage.DayStatsWithTimestamp{}
}
func (s *SimpleStorage) GetAddressStats(addr string) *storage.AddressStats {
	return &storage.AddressStats{}
}
func (s *SimpleStorage) GenesisHashExist() bool              { return false }
func (s *SimpleStorage) GetGenesisHash() storage.GenesisHash { return storage.GenesisHash("") }
func (s *SimpleStorage) GetTransactionByHash(hash string) models.Transaction {
	return models.Transaction{}
}
func (s *SimpleStorage) GetInternalTransactions(hash string) models.InternalTransactionsResponse {
	return models.InternalTransactionsResponse{}
}
func (s *SimpleStorage) GetTransactionLogs(hash string) models.TransactionLogsResponse {
	return models.TransactionLogsResponse{}
}
func (s *SimpleStorage) GetValidatorYield(validator string, block uint64) storage.Yield {
	return storage.Yield{}
}
func (s *SimpleStorage) GetTotalYield(block uint64) storage.Yield         { return storage.Yield{} }
func (s *SimpleStorage) GetMonthlyActiveAddresses(to_date uint64) *uint64 { return nil }
func (s *SimpleStorage) GetDailyContractUsers(address string, timestamp uint64) storage.DailyContractUsersList {
	return storage.DailyContractUsersList{}
}

// Test error handling in generated server code
func TestGeneratedServerErrorHandling(t *testing.T) {
	e := echo.New()

	// Setup error handler (same as main.go)
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = ctx.JSON(he.Code, map[string]any{"message": he.Message})
			return
		}
		errMsg := err.Error()
		_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": errMsg})
	}

	// Create handler
	apiHandler := &ApiHandler{
		storage: &SimpleStorage{},
		config:  common.DefaultConfig(),
		stats:   nil,
	}

	// Register handlers
	RegisterHandlers(e, apiHandler)

	// Test invalid parameter
	req := httptest.NewRequest(http.MethodGet, "/contractStats?fromDate=invalid&toDate=1735689600", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	// Should return 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response, "message")
}

// Test that empty arrays are properly JSON serialized
func TestEmptyArraySerialization(t *testing.T) {
	// Test that our response type correctly serializes empty arrays
	response := models.ContractsResponse{
		Contracts: make([]models.Address, 0),
	}

	data, err := json.Marshal(response)
	assert.NoError(t, err)

	// Should contain empty array, not null
	assert.Contains(t, string(data), `"contracts":[]`)
	assert.NotContains(t, string(data), `"contracts":null`)

	// Test with nil slice (which might happen in some cases)
	responseWithNil := models.ContractsResponse{
		Contracts: nil,
	}

	dataNil, err := json.Marshal(responseWithNil)
	assert.NoError(t, err)

	// This will be null, which is why we need to initialize with make()
	assert.Contains(t, string(dataNil), `"contracts":null`)
}
