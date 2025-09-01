package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareTestSuite struct {
	suite.Suite
	server *echo.Echo
}

func (suite *AuthMiddlewareTestSuite) SetupTest() {
	e := echo.New()

	// Create a minimal OpenAPI spec for testing
	swagger := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   "Test API",
			Version: "1.0.0",
		},
		Paths: &openapi3.Paths{},
	}

	// Setup authentication
	SetupOpenAPIAuth(e, swagger, Config{
		Username: "testuser",
		Password: "testpass",
	})

	// Add test routes
	e.GET("/public", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "public"})
	})

	e.GET("/contracts", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "protected"})
	})

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

func (suite *AuthMiddlewareTestSuite) TestPublicEndpoint() {
	req := httptest.NewRequest(http.MethodGet, "/public", nil)
	rec := httptest.NewRecorder()

	suite.server.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "public")
}

func (suite *AuthMiddlewareTestSuite) TestProtectedEndpointWithValidAuth() {
	req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
	req.SetBasicAuth("testuser", "testpass")
	rec := httptest.NewRecorder()

	suite.server.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "protected")
}

func (suite *AuthMiddlewareTestSuite) TestProtectedEndpointWithoutAuth() {
	req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
	rec := httptest.NewRecorder()

	suite.server.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "Missing Authorization header")
}

func (suite *AuthMiddlewareTestSuite) TestProtectedEndpointWithInvalidAuth() {
	req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
	req.SetBasicAuth("wrong", "credentials")
	rec := httptest.NewRecorder()

	suite.server.ServeHTTP(rec, req)

	assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "Invalid credentials")
}

func (suite *AuthMiddlewareTestSuite) TestInvalidAuthorizationHeader() {
	testCases := []struct {
		name        string
		authHeader  string
		expectedMsg string
	}{
		{
			name:        "Bearer token instead of Basic",
			authHeader:  "Bearer token123",
			expectedMsg: "Invalid Authorization header format",
		},
		{
			name:        "Invalid base64",
			authHeader:  "Basic !!invalid!!",
			expectedMsg: "Invalid base64 encoding",
		},
		{
			name:        "Missing colon in credentials",
			authHeader:  "Basic dGVzdA==", // "test" without colon
			expectedMsg: "Invalid credentials format",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
			req.Header.Set("Authorization", tc.authHeader)
			rec := httptest.NewRecorder()

			suite.server.ServeHTTP(rec, req)

			assert.Equal(suite.T(), http.StatusUnauthorized, rec.Code)
			assert.Contains(suite.T(), rec.Body.String(), tc.expectedMsg)
		})
	}
}

func TestAuthMiddlewareSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}

func TestSetupOpenAPIAuthWithoutCredentials(t *testing.T) {
	e := echo.New()
	swagger := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:   "Test API",
			Version: "1.0.0",
		},
		Paths: &openapi3.Paths{},
	}

	// Setup without credentials
	SetupOpenAPIAuth(e, swagger, Config{})

	// Add test route that would normally be protected
	e.GET("/contracts", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "accessible"})
	})

	// Test that endpoint is accessible without auth when no credentials configured
	req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "accessible")
}

func TestCreateAuthMiddleware(t *testing.T) {
	// Test the middleware function directly
	middleware := createAuthMiddleware("user", "pass")

	// Create a mock handler
	handler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "success"})
	}

	// Wrap handler with middleware
	wrappedHandler := middleware(handler)

	testCases := []struct {
		name           string
		path           string
		setupAuth      func(*http.Request)
		expectedStatus int
	}{
		{
			name: "Public endpoint",
			path: "/public",
			setupAuth: func(req *http.Request) {
				// No auth needed
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Protected endpoint with valid auth",
			path: "/contracts",
			setupAuth: func(req *http.Request) {
				req.SetBasicAuth("user", "pass")
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "Protected endpoint without auth",
			path: "/contracts",
			setupAuth: func(req *http.Request) {
				// No auth
			},
			expectedStatus: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			e.HTTPErrorHandler = func(err error, ctx echo.Context) {
				if he, ok := err.(*echo.HTTPError); ok {
					_ = ctx.JSON(he.Code, map[string]any{"message": he.Message})
					return
				}
				_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": err.Error()})
			}

			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			tc.setupAuth(req)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := wrappedHandler(c)
			if err != nil {
				e.HTTPErrorHandler(err, c)
			}

			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}

// Benchmark authentication middleware
func BenchmarkAuthMiddleware(b *testing.B) {
	middleware := createAuthMiddleware("benchuser", "benchpass")
	handler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
	}
	wrappedHandler := middleware(handler)

	e := echo.New()
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = ctx.JSON(he.Code, map[string]any{"message": he.Message})
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
		req.SetBasicAuth("benchuser", "benchpass")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := wrappedHandler(c)
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}
	}
}

func BenchmarkAuthMiddleware_InvalidCredentials(b *testing.B) {
	middleware := createAuthMiddleware("benchuser", "benchpass")
	handler := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
	}
	wrappedHandler := middleware(handler)

	e := echo.New()
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			_ = ctx.JSON(he.Code, map[string]any{"message": he.Message})
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/contracts", nil)
		req.SetBasicAuth("wrong", "credentials")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := wrappedHandler(c)
		if err == nil {
			b.Errorf("Expected authentication error but got none")
		}
	}
}
