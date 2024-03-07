package controllers

import (
	"net/http"
	"scm-api/services/operations_service"
	operation_types "scm-api/types/operations/requests"

	"github.com/labstack/echo/v4"
)

func CreateOperation(c echo.Context) error {
	req := c.Get("validatedRequest").(*operation_types.CreateOperationRequest)

	if operations_service.OperationExists(req.Name) {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Operation with same name already exists"})
	}

	newOperation, err := operations_service.CreateOperation(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create operation"})
	}

	return c.JSON(http.StatusCreated, newOperation)
}
