package controllers

import (
	"net/http"
	"scm-api/services/vehicles_service"
	vehicle_requests "scm-api/types/vehicles/requests"

	"github.com/labstack/echo/v4"
)

func CreateVehicle(c echo.Context) error {
	req := c.Get("validatedRequest").(*vehicle_requests.CreateVehicleRequest)

	if vehicles_service.VehicleExists(req.Name) {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Vehicle with same name already exists"})
	}

	err := vehicles_service.ValidateAttributes(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newVehicle, err := vehicles_service.CreateVehicle(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create vehicle"})
	}

	return c.JSON(http.StatusCreated, newVehicle)
}

func GetVehicles(c echo.Context) error {
	vehicles, error := vehicles_service.GetAllVehicles()

	if error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve vehicles"})
	}

	return c.JSON(http.StatusOK, vehicles)
}
