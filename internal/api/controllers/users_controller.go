package controllers

import (
	"net/http"

	dtos "scm-api/internal/api/dtos/requests"
	"scm-api/internal/services/users_service"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	req := c.Get("validatedRequest").(*dtos.CreateUserRequest)

	if users_service.UserExists(req.Username) {
		return c.JSON(http.StatusConflict, map[string]string{"error": "User already exists"})
	}

	newUser, err := users_service.CreateUser(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, newUser)
}
