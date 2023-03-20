package main

import (
	"github.com/Calgorr/URL_Shortener/handle"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/:hash", handle.Redirect)
	e.POST("/new", handle.SaveUrl)
}
