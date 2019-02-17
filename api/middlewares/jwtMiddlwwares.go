package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJWTMiddlewares(g *echo.Group) {
	// /jwt にリクエストを送った際にJWTトークンでの認証を行う
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))
}
