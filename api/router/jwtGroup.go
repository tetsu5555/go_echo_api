package router

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func JWTGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
