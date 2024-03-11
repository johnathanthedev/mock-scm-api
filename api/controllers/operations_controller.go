package controllers

import (
	"errors"
	"net/http"
	"scm-api/services/operations_service"
	"scm-api/services/users_service"
	"scm-api/services/vehicles_service"

	operation_requests "scm-api/types/operations/requests"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOperation(c echo.Context) error {
	req := c.Get("validatedRequest").(*operation_requests.CreateOperationRequest)

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
	req, ok := c.Get("validatedRequest").(*operation_requests.JoinOperationRequest)

	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
	}

	username := c.Request().Header.Get("Authorization")

	user_id, err := users_service.GetUserIdByUsername(username)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user ID"})
		}
	}

	userAlreadyInOperation, err := operations_service.IsUserInOperation(req.OperationID, user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check operation membership"})
	}

	err = operations_service.AddUserToOperation(req.OperationID, user_id)
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

func ListOperations(c echo.Context) error {
	operations, err := operations_service.ListOperations()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve operations"})
	}

	return c.JSON(http.StatusOK, operations)
}

func ListUserJoinedOperations(c echo.Context) error {
	username := c.Request().Header.Get("Authorization")

	user_id, err := users_service.GetUserIdByUsername(username)
	if err != nil {
		if err.Error() == "user not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user ID"})
		}
	}

	operations, err := operations_service.ListUserJoinedOperations(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve operations"})
	}

	return c.JSON(http.StatusOK, operations)
}

func AddVehicleToOperation(c echo.Context) error {
	req, _ := c.Get("validatedRequest").(*operation_requests.AddVehicleToOperationRequest)

	operation, err := operations_service.GetOperationByID(req.OperationID)
	if err != nil {
		if err.Error() == "operation not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Operation not found"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check operation existence"})
		}
	}

	vehicle, err := vehicles_service.GetVehicleById(req.VehicleID)
	if err != nil {
		if err.Error() == "vehicle not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Vehicle not found"})
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check vehicle existence"})
		}
	}

	vehicleAlreadyInOperation := operations_service.IsVehicleInOperation(operation.ID, vehicle.ID)

	if vehicleAlreadyInOperation {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Vehicle is already part of the operation"})
	}

	err = operations_service.AddVehicleToOperation(operation, vehicle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add vehicle to operation"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Vehicle successfully added to operation"})
}
