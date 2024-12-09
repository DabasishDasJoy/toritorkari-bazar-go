package main

import (
	"toritorkari-bazar/cmd"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cmd.Serve(e)
}
