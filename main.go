package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.File("/", "./textfiles/dota_fan_page.html")
	e.Static("/", "./textfiles")
	e.Static("/img/", "./img")
	e.Logger.Fatal(e.Start(":8080"))
}
