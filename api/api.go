package api

import (
	"os"
	"scm-api/api/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() error {
	e := echo.New()

	frontend_url := os.Getenv("FRONTEND_URL")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontend_url},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	cv := validator.Init()

	InitRoutes(e, cv)

	return e.Start(":8080")
}
