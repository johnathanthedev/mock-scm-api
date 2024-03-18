package controllers

import (
	"net/http"
	"scm-api/services/facilities_service"
	"scm-api/services/route_stops_service"
	"scm-api/services/routes_service"
	route_stop_dtos "scm-api/types/route-stops/dtos"

	"github.com/labstack/echo/v4"
)

func CreateRouteStop(c echo.Context) error {
	req := c.Get("validatedRequest").(*route_stop_dtos.CreateRouteStopDto)

	_, route_err := routes_service.GetRouteByID(req.RouteID)
	if route_err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Route not found"})
	}

	_, facility_err := facilities_service.GetFacilityByID(req.FacilityID)
	if facility_err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Facility not found"})
	}

	newRouteStop, err := route_stops_service.CreateRouteStop(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create route stop"})
	}

	return c.JSON(http.StatusCreated, newRouteStop)
}

func ListRouteStops(c echo.Context) error {
	req := c.Get("validatedRequest").(*route_stop_dtos.ListRouteStopsDto)

	_, route_err := routes_service.GetRouteByID(req.RouteID)
	if route_err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Route not found"})
	}

	stops, err := route_stops_service.GetRouteStopsByRouteID(req.RouteID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve route stops"})
	}

	return c.JSON(http.StatusOK, stops)
}
