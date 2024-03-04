package api

import (
	controllers "scm-api/internal/api/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
    e.GET("/users", controllers.UsersController)
}
