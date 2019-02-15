package main

import (
	"net/http"

	"./animal"
	"./cat"
	"./handler"
	"./interceptor"
	"./template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 静的ファイルのパスを設定
	e.Static("/static", "static")

	// ルーティング
	e.GET("/hello/:username", handler.MainPage(), interceptor.BasicAuth())

	e.GET("/cats/:data", cat.GetCat)

	e.POST("/cats", animal.AddCat)
	e.POST("/dogs", animal.AddDog)
	e.POST("/hamsters", animal.AddHamster)

	e.Renderer = template.Renderer
	// Named route "foobar"
	e.GET("/template", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"
	e.Logger.Fatal(e.Start(":8000"))

	// サーバー起動
	e.Start(":8000") //ポート番号指定してね
}
