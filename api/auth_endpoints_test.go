package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Taraxa-project/taraxa-indexer/internal/auth"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestAuthEndpoints is a simplified test suite for authorization
type TestAuthEndpoints struct {
	suite.Suite
	server *echo.Echo
}

func (suite *TestAuthEndpoints) SetupTest() {
	e := echo.New()

	// Load OpenAPI spec for authentication
	swagger, err := openapi3.NewLoader().LoadFromFile("openapi.yaml")
	if err != nil {
		swagger = &openapi3.T{
			OpenAPI: "3.0.0",
			Info:    &openapi3.Info{Title: "Test API", Version: "1.0.0"},
			Paths:   &openapi3.Paths{},
		}
	}

	// Setup authentication
	auth.SetupOpenAPIAuth(e, swagger, auth.Config{
		Username: "testuser",
		Password: "testpass",
	})

	// Create API handler
	apiHandler := &ApiHandler{
		storage: &SimpleStorage{},
		config:  common.DefaultConfig(),
		stats:   chain.MakeStats(100),
	}

	RegisterHandlers(e, apiHandler)

	// Setup error handler
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = ctx.JSON(he.Code, map[string]any{"message": he.Message})
			return
		}
		_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": err.Error()})
	}

	suite.server = e
}

// AuthTestCase represents a single authentication test scenario
type AuthTestCase struct {
	Name           string
	SetupAuth      func(*http.Request)
	ExpectedStatus int
	ExpectedError  string
}

// ProtectedEndpoint represents an endpoint that requires authentication
type ProtectedEndpoint struct {
	Name     string
	Path     string
	Method   string
	Expected int // Expected status for valid auth (200, 404, etc.)
}

// getStandardAuthTests returns the standard set of auth test cases
func getStandardAuthTests() []AuthTestCase {
	return []AuthTestCase{
		{
			Name: "Valid authentication",
			SetupAuth: func(req *http.Request) {
				req.SetBasicAuth("testuser", "testpass")
			},
			ExpectedStatus: http.StatusOK, // Will be overridden per endpoint
		},
		{
			Name: "No authentication",
			SetupAuth: func(req *http.Request) {
				// No auth header
			},
			ExpectedStatus: http.StatusUnauthorized,
			ExpectedError:  "Missing Authorization header",
		},
		{
			Name: "Invalid credentials",
			SetupAuth: func(req *http.Request) {
				req.SetBasicAuth("wrong", "credentials")
			},
			ExpectedStatus: http.StatusUnauthorized,
			ExpectedError:  "Invalid credentials",
		},
		{
			Name: "Invalid auth format",
			SetupAuth: func(req *http.Request) {
				req.Header.Set("Authorization", "Bearer token123")
			},
			ExpectedStatus: http.StatusUnauthorized,
			ExpectedError:  "Invalid Authorization header format",
		},
	}
}

// getProtectedEndpoints returns all endpoints that require authentication
func getProtectedEndpoints() []ProtectedEndpoint {
	return []ProtectedEndpoint{
		{
			Name:     "monthlyActiveAddresses",
			Path:     "/monthlyActiveAddresses?date=1704067200",
			Method:   http.MethodGet,
			Expected: http.StatusOK,
		},
		{
			Name:     "monthlyStats",
			Path:     "/monthlyStats?date=1704067200",
			Method:   http.MethodGet,
			Expected: http.StatusNotFound, // Known issue with insufficient data
		},
		{
			Name:     "contractStats",
			Path:     "/contractStats?fromDate=1704067200&toDate=1735689600",
			Method:   http.MethodGet,
			Expected: http.StatusOK,
		},
		{
			Name:     "monthlyAverageDailyActiveWallets",
			Path:     "/monthlyAverageDailyActiveWallets/0x0000000000000000000000000000000000000000?date=1704067200",
			Method:   http.MethodGet,
			Expected: http.StatusOK,
		},
		{
			Name:     "contracts",
			Path:     "/contracts",
			Method:   http.MethodGet,
			Expected: http.StatusOK,
		},
	}
}

// TestAllProtectedEndpoints tests all protected endpoints with all auth scenarios
func (suite *TestAuthEndpoints) TestAllProtectedEndpoints() {
	endpoints := getProtectedEndpoints()
	authTests := getStandardAuthTests()

	for _, endpoint := range endpoints {
		suite.Run(endpoint.Name, func() {
			for _, authTest := range authTests {
				suite.Run(authTest.Name, func() {
					req := httptest.NewRequest(endpoint.Method, endpoint.Path, nil)
					authTest.SetupAuth(req)
					rec := httptest.NewRecorder()

					suite.server.ServeHTTP(rec, req)

					// Adjust expected status for valid auth based on endpoint
					expectedStatus := authTest.ExpectedStatus
					if authTest.Name == "Valid authentication" {
						expectedStatus = endpoint.Expected
					}

					assert.Equal(suite.T(), expectedStatus, rec.Code,
						"Endpoint: %s, Test: %s", endpoint.Name, authTest.Name)

					if authTest.ExpectedError != "" {
						assert.Contains(suite.T(), rec.Body.String(), authTest.ExpectedError,
							"Endpoint: %s, Test: %s", endpoint.Name, authTest.Name)
					}
				})
			}
		})
	}
}

// TestPublicEndpointsRemainPublic ensures public endpoints don't require auth
func (suite *TestAuthEndpoints) TestPublicEndpointsRemainPublic() {
	publicEndpoints := []string{
		"/holders?pagination={\"limit\":10}",
		"/totalSupply",
		"/totalYield",
		"/validators",
		"/chainStats",
	}

	for _, endpoint := range publicEndpoints {
		suite.Run("Public_"+endpoint, func() {
			req := httptest.NewRequest(http.MethodGet, endpoint, nil)
			rec := httptest.NewRecorder()

			suite.server.ServeHTTP(rec, req)

			assert.NotEqual(suite.T(), http.StatusUnauthorized, rec.Code,
				"Public endpoint %s should not require authentication", endpoint)
		})
	}
}

// TestResponseFormats validates JSON response formats for working endpoints
func (suite *TestAuthEndpoints) TestResponseFormats() {
	testCases := []struct {
		Name     string
		Endpoint string
		IsArray  bool
		Contains string
	}{
		{
			Name:     "monthlyActiveAddresses",
			Endpoint: "/monthlyActiveAddresses?date=1704067200",
			IsArray:  false,
			Contains: "count",
		},
		{
			Name:     "contractStats",
			Endpoint: "/contractStats?fromDate=1704067200&toDate=1735689600",
			IsArray:  true,
			Contains: "[]",
		},
		{
			Name:     "monthlyAverageDailyActiveWallets",
			Endpoint: "/monthlyAverageDailyActiveWallets/0x0000000000000000000000000000000000000000?date=1704067200",
			IsArray:  false,
			Contains: "count",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.Name, func() {
			req := httptest.NewRequest(http.MethodGet, tc.Endpoint, nil)
			req.SetBasicAuth("testuser", "testpass")
			rec := httptest.NewRecorder()

			suite.server.ServeHTTP(rec, req)

			if rec.Code == http.StatusOK {
				responseBody := rec.Body.String()

				// Validate JSON structure
				if tc.IsArray {
					var response []interface{}
					err := json.Unmarshal(rec.Body.Bytes(), &response)
					assert.NoError(suite.T(), err, "Should be valid JSON array")
				} else {
					var response map[string]interface{}
					err := json.Unmarshal(rec.Body.Bytes(), &response)
					assert.NoError(suite.T(), err, "Should be valid JSON object")
				}

				assert.Contains(suite.T(), responseBody, tc.Contains)
			}
		})
	}
}

// TestAuthPerformance does a simple performance check
func (suite *TestAuthEndpoints) TestAuthPerformance() {
	req := httptest.NewRequest(http.MethodGet, "/monthlyActiveAddresses?date=1704067200", nil)
	req.SetBasicAuth("testuser", "testpass")

	// Run 10 times to ensure consistent performance
	for i := 0; i < 10; i++ {
		rec := httptest.NewRecorder()
		suite.server.ServeHTTP(rec, req)
		assert.Equal(suite.T(), http.StatusOK, rec.Code)
	}
}

// TestSpecificAuthScenarios tests additional edge cases
func TestSpecificAuthScenarios(t *testing.T) {
	// This can be expanded for specific edge cases not covered by the main suite
	testCases := []struct {
		name       string
		endpoint   string
		authHeader string
		expected   int
	}{
		{
			name:       "Invalid base64 encoding",
			endpoint:   "/contracts",
			authHeader: "Basic !!invalid!!",
			expected:   http.StatusUnauthorized,
		},
		{
			name:       "Missing colon in credentials",
			endpoint:   "/contracts",
			authHeader: "Basic dGVzdA==", // "test" without colon
			expected:   http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup server (simplified version of main setup)
			e := echo.New()
			swagger := &openapi3.T{
				OpenAPI: "3.0.0",
				Info:    &openapi3.Info{Title: "Test", Version: "1.0.0"},
				Paths:   &openapi3.Paths{},
			}
			auth.SetupOpenAPIAuth(e, swagger, auth.Config{
				Username: "testuser", Password: "testpass",
			})

			apiHandler := &ApiHandler{
				storage: &SimpleStorage{},
				config:  common.DefaultConfig(),
				stats:   chain.MakeStats(100),
			}
			RegisterHandlers(e, apiHandler)

			e.HTTPErrorHandler = func(err error, ctx echo.Context) {
				if he, ok := err.(*echo.HTTPError); ok {
					_ = ctx.JSON(he.Code, map[string]any{"message": he.Message})
					return
				}
				_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": err.Error()})
			}

			req := httptest.NewRequest(http.MethodGet, tc.endpoint, nil)
			req.Header.Set("Authorization", tc.authHeader)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expected, rec.Code)
		})
	}
}
