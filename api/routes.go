package api

import (
	controllers "scm-api/api/controllers"
	vehicle_controllers "scm-api/api/controllers/vehicles"
	facilities_dto "scm-api/types/facilities/dtos"
	operation_types "scm-api/types/operations/requests"
	product_dtos "scm-api/types/products/dtos"
	route_dtos "scm-api/types/routes/dtos"
	user_requests "scm-api/types/users/requests"
	vehicle_types "scm-api/types/vehicles/requests"

	"scm-api/api/middleware"
	"scm-api/api/validator"

	ws "scm-api/ws"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, cv *validator.CustomValidator, broker *ws.Broker) {
	// ===================================================
	// WS
	// ===================================================
	e.GET(
		"/ws",
		controllers.WebSocketHandler(broker),
		// middleware.AuthorizationMiddleware(),
	)

	// ===================================================
	// Users
	// ===================================================
	e.POST("/users/create", controllers.CreateUser, middleware.ValidationsMiddleware(cv, &user_requests.CreateUserRequest{}))
	e.POST("/users/login", controllers.Login, middleware.ValidationsMiddleware(cv, &user_requests.LoginRequest{}))

	// ===================================================
	// Operations
	// ===================================================
	e.POST(
		"/operations/create",
		controllers.CreateOperation,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &operation_types.CreateOperationRequest{}),
	)
	e.GET("/operations/list", controllers.ListOperations, middleware.AuthorizationMiddleware())
	e.GET("/operations/user-joined", controllers.ListUserJoinedOperations, middleware.AuthorizationMiddleware())
	e.POST(
		"/operations/join",
		controllers.JoinOperation,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &operation_types.JoinOperationRequest{}),
	)
	e.POST(
		"/operations/vehicles/add-to-operation",
		controllers.AddVehicleToOperation,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &operation_types.AddVehicleToOperationRequest{}),
	)

	// ===================================================
	// Facilities
	// ===================================================
	e.POST(
		"/facilities/create",
		controllers.CreateFacility,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &facilities_dto.CreateFacilityDto{}),
	)

	// ===================================================
	// Vehicles
	// ===================================================
	vehicleCtrl := vehicle_controllers.NewVehicleTrackingController(broker)

	e.GET("/vehicles/list", vehicle_controllers.GetVehicles, middleware.AuthorizationMiddleware())
	e.POST(
		"/vehicles/create",
		vehicle_controllers.CreateVehicle,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &vehicle_types.CreateVehicleRequest{}),
	)
	e.POST("/vehicles/update-location", vehicleCtrl.UpdateVehicleLocation, middleware.ValidationsMiddleware(cv, &vehicle_types.UpdateVehicleLocationRequest{}))

	// ===================================================
	// Products
	// ===================================================
	e.POST(
		"/products/create",
		controllers.CreateProduct,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &product_dtos.CreateProductDto{}),
	)
	e.POST(
		"/products/list",
		controllers.ListProducts,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &product_dtos.ListProductsDto{}),
	)

	// ===================================================
	// Routes
	// ===================================================
	e.POST(
		"/routes/create",
		controllers.CreateRoute,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &route_dtos.CreateRouteDto{}),
	)
	e.POST(
		"/routes/list",
		controllers.ListRoutes,
		middleware.AuthorizationMiddleware(),
		middleware.ValidationsMiddleware(cv, &route_dtos.ListRoutesDto{}),
	)
}
