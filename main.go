package main

import (
	"net/http"
	"toritorkari-bazar/cmd"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	cmd.Serve(e)
}
