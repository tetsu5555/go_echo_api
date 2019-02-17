package api

import (
	"./handlers"

	"github.com/labstack/echo"
)

func AdminGroup(g *echo.Group) {
	// /admin/maiにリクエストした際に、mainAdminが呼び出されるようになる
	g.GET("/main", handlers.MainAdmin)
}
