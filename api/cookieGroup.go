package api

import (
	"./handlers"

	"github.com/labstack/echo"
)

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie)
}
