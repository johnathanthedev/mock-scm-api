package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UsersController(c echo.Context) error {
    users := []string{"user1", "user2", "user3"}
    return c.JSON(http.StatusOK, users)
}
