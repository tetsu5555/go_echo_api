package router

import (
	"api/middlewares"

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
	MainGroup(e)

	// set groups routes
	AdminGroup(adminGroup)
	CookieGroup(cookieGroup)
	JWTGroup(jwtGroup)

	return e
}
