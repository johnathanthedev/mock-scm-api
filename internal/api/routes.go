package api

import (
	controllers "scm-api/internal/api/controllers"

	"scm-api/internal/api/middleware"
	"scm-api/internal/api/validator"

	dtos "scm-api/internal/api/dtos/requests"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, cv *validator.CustomValidator) {
	e.POST("/users/create", controllers.CreateUser, middleware.ValidationsMiddleware(cv, &dtos.CreateUserRequest{}))
	e.POST("/users/login", controllers.Login, middleware.ValidationsMiddleware(cv, &dtos.LoginRequest{}))
}
