package api

import (
	controllers "scm-api/api/controllers"

	"scm-api/api/middleware"
	"scm-api/api/validator"

	dtos "scm-api/api/dtos/requests"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, cv *validator.CustomValidator) {
	e.POST("/users/create", controllers.CreateUser, middleware.ValidationsMiddleware(cv, &dtos.CreateUserRequest{}))
	e.POST("/users/login", controllers.Login, middleware.ValidationsMiddleware(cv, &dtos.LoginRequest{}))
}
