package controllers

import (
	"net/http"

	dtos "scm-api/api/dtos/requests"
	"scm-api/services/users_service"

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

func Login(c echo.Context) error {
	req := c.Get("validatedRequest").(*dtos.LoginRequest)

	if !users_service.UserExists(req.Username) {
		return c.JSON(http.StatusConflict, map[string]string{"error": "User not found"})
	}

	response_message := map[string]string{"message": "Login successful"}

	return c.JSON(http.StatusOK, response_message)
}
