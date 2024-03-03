package api

import (
	"github.com/labstack/echo/v4"
)

func StartServer() error {
    e := echo.New()

    InitRoutes(e)

    return e.Start(":8080")
}
