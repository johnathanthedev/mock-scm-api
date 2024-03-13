package middleware

import (
	"net/http"
	"scm-api/services/users_service"

	"github.com/labstack/echo/v4"
)

// AuthorizationMiddleware checks if the user is authorized
func AuthorizationMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			username := c.Request().Header.Get("Authorization")

			if username == "" {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Authorization header is empty"})
			}

			if !users_service.UserExists(username) {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Access denied"})
			}

			return next(c)
		}
	}
}
