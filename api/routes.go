package api

import (
	controllers "scm-api/api/controllers"
	operation_types "scm-api/types/operations/requests"
	user_requests "scm-api/types/users/requests"
	vehicle_types "scm-api/types/vehicles/requests"

	"scm-api/api/middleware"
	"scm-api/api/validator"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, cv *validator.CustomValidator) {
	e.POST("/users/create", controllers.CreateUser, middleware.ValidationsMiddleware(cv, &user_requests.CreateUserRequest{}))
	e.POST("/users/login", controllers.Login, middleware.ValidationsMiddleware(cv, &user_requests.LoginRequest{}))

	e.POST("/operations/create", controllers.CreateOperation, middleware.ValidationsMiddleware(cv, &operation_types.CreateOperationRequest{}))

	e.GET("/vehicles/list", controllers.GetVehicles)
	e.POST("/vehicles/create", controllers.CreateVehicle, middleware.ValidationsMiddleware(cv, &vehicle_types.CreateVehicleRequest{}))
}
