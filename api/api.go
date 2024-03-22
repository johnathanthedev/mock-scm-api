package api

import (
	"os"
	"scm-api/api/validator"

	ws "scm-api/ws"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer(broker *ws.Broker) error {
	e := echo.New()

	frontend_url := os.Getenv("FRONTEND_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{frontend_url},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	cv := validator.Init()

	InitRoutes(e, cv, broker)

	return e.Start(":" + port)
}
