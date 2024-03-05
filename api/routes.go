package api

import (
	controllers "scm-api/api/controllers"
	dtos "scm-api/api/dtos/requests"

	"scm-api/api/middleware"
	"scm-api/api/validator"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, cv *validator.CustomValidator) {
	e.POST("/users/create", controllers.CreateUser, middleware.ValidationsMiddleware(cv, &dtos.CreateUserRequest{}))
	e.POST("/users/login", controllers.Login, middleware.ValidationsMiddleware(cv, &dtos.LoginRequest{}))

	e.POST("/operations/create", controllers.CreateOperation, middleware.ValidationsMiddleware(cv, &dtos.CreateOperationRequest{}))
}
