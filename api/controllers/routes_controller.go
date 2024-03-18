package controllers

import (
	"net/http"
	"scm-api/services/facilities_service"
	"scm-api/services/operations_service"
	"scm-api/services/routes_service"
	route_dtos "scm-api/types/routes/dtos"

	"github.com/labstack/echo/v4"
)

func CreateRoute(c echo.Context) error {
	req := c.Get("validatedRequest").(*route_dtos.CreateRouteDto)

	_, operation_err := operations_service.GetOperationByID(req.OperationID)
	if operation_err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
	}

	_, facility_err := facilities_service.GetFacilityByID(req.OriginFacilityID)
	if facility_err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Facility not found"})
	}

	newRoute, err := routes_service.CreateRoute(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create Route"})
	}

	return c.JSON(http.StatusCreated, newRoute)
}

func ListRoutes(c echo.Context) error {
	req := c.Get("validatedRequest").(*route_dtos.ListRoutesDto)

	operation_id := req.OperationID

	_, err := operations_service.GetOperationByID(operation_id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
	}

	routes, err := routes_service.GetAllRoutesByOperationID(operation_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve routes"})
	}

	return c.JSON(http.StatusOK, routes)
}
