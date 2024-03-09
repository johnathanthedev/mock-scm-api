package controllers

import (
	"errors"
	"net/http"
	"scm-api/services/operations_service"
	operation_types "scm-api/types/operations/requests"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func JoinOperation(c echo.Context) error {
	req, ok := c.Get("validatedRequest").(*operation_types.JoinOperationRequest)

	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	userAlreadyInOperation, err := operations_service.UserInOperation(req.OperationID, req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check operation membership"})
	}

	err = operations_service.AddUserToOperation(req.OperationID, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to join operation"})
	}

	if userAlreadyInOperation {
		return c.JSON(http.StatusConflict, map[string]string{"error": "User is already part of the operation"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User successfully joined operation"})
}
