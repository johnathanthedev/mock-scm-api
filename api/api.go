package api

import (
	"scm-api/api/validator"

	"github.com/labstack/echo/v4"
)

func StartServer() error {
	e := echo.New()

	cv := validator.Init()

	InitRoutes(e, cv)

	return e.Start(":8080")
}
