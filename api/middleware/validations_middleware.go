package middleware

import (
	"net/http"
	"reflect"

	"scm-api/api/validator"

	"github.com/labstack/echo/v4"
)

// ValidationsMiddleware validates incoming requests using CustomValidator.
func ValidationsMiddleware(cv *validator.CustomValidator, reqType interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			req := reflect.New(reflect.TypeOf(reqType).Elem()).Interface()
			if err := c.Bind(req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}

			if err := cv.Validate(req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			}

			// Set the validated request in the context
			c.Set("validatedRequest", req)

			return next(c)
		}
	}
}
