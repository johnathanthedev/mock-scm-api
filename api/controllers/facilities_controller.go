package controllers

import (
	"net/http"
	"scm-api/services/facilities_service"
	"scm-api/services/operations_service"
	facility_dtos "scm-api/types/facilities/dtos"

	"github.com/labstack/echo/v4"
)

func CreateFacility(c echo.Context) error {
	req := c.Get("validatedRequest").(*facility_dtos.CreateFacilityDto)

	newOperation, err := facilities_service.CreateFacility(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create operation"})
	}

	return c.JSON(http.StatusCreated, newOperation)
}

func ListOperationFacilities(c echo.Context) error {
	req := c.Get("validatedRequest").(*facility_dtos.ListFacilitiesDto)

	_, err := operations_service.GetOperationByID(req.OperationID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
	}

	facilities, err := facilities_service.GetAllFacilitiesByOperationID(req.OperationID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve facilities"})
	}

	return c.JSON(http.StatusCreated, facilities)
}
