package auth

import (
	"crypto/subtle"
	"encoding/base64"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// Config holds authentication configuration
type Config struct {
	Username string
	Password string
}

// SetupOpenAPIAuth configures authentication middleware without OpenAPI validation
func SetupOpenAPIAuth(e *echo.Echo, swagger *openapi3.T, config Config) {
	// Skip OpenAPI validation middleware entirely to avoid status code issues
	// Use custom middleware that only handles authentication for protected endpoints

	if config.Username != "" && config.Password != "" {
		log.WithField("username", config.Username).Info("Configured custom authentication middleware for Basic Auth")
		e.Use(createAuthMiddleware(config.Username, config.Password))
	} else {
		log.Info("No authentication configured - all endpoints are public")
	}
}

// createAuthMiddleware creates a custom authentication middleware
func createAuthMiddleware(username, password string) echo.MiddlewareFunc {
	// Define which endpoints require authentication
	protectedPaths := map[string]bool{
		"/contracts":              true,
		"/contractStats":          true,
		"/monthlyActiveAddresses": true,
		"/monthlyStats":           true,
	}

	// Also check for pattern matches (for parametrized paths)
	protectedPatterns := []string{
		"/monthlyAverageDailyActiveWallets/",
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Check if this endpoint requires authentication
			path := c.Request().URL.Path
			requiresAuth := protectedPaths[path]

			// Check pattern matches for parametrized paths
			if !requiresAuth {
				for _, pattern := range protectedPatterns {
					if strings.HasPrefix(path, pattern) {
						requiresAuth = true
						break
					}
				}
			}

			if requiresAuth {
				log.Debug("Endpoint requires authentication")

				// Get the Authorization header
				authHeader := c.Request().Header.Get("Authorization")
				if authHeader == "" {
					return echo.NewHTTPError(401, "Missing Authorization header")
				}

				// Parse Basic Auth
				const prefix = "Basic "
				if !strings.HasPrefix(authHeader, prefix) {
					return echo.NewHTTPError(401, "Invalid Authorization header format")
				}

				// Decode base64
				encoded := authHeader[len(prefix):]
				decoded, err := base64.StdEncoding.DecodeString(encoded)
				if err != nil {
					return echo.NewHTTPError(401, "Invalid base64 encoding")
				}

				// Split username:password
				credentials := string(decoded)
				parts := strings.SplitN(credentials, ":", 2)
				if len(parts) != 2 {
					return echo.NewHTTPError(401, "Invalid credentials format")
				}

				user, pass := parts[0], parts[1]

				// Validate credentials using constant-time comparison to prevent timing attacks
				if subtle.ConstantTimeCompare([]byte(user), []byte(username)) == 1 &&
					subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1 {
					log.WithField("username", user).Info("Authentication successful for protected endpoint")
					return next(c) // Authentication successful, continue
				}

				return echo.NewHTTPError(401, "Invalid credentials")
			}

			// No authentication required for this endpoint
			return next(c)
		}
	}
}
