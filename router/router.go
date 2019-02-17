package router

import (
	"../api"
	"../api/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// Echoのインスタンス作る
	e := echo.New()

	// Groupを作成
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set middleware
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJWTMiddlewares(jwtGroup)

	// set main routes
	api.MainGroup(e)

	// set groups routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	api.JWTGroup(jwtGroup)

	return e
}
