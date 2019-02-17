package api

import (
	"./handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	// ルーティング
	// 第三引数に渡したミドルウェアでBasicAuthenticationを行なっている
	e.GET("/hello/:username", handlers.MainPage(), handlers.BasicAuth())

	e.GET("/cats/:data", handlers.GetCat)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
	e.GET("/login", handlers.Login)
}
