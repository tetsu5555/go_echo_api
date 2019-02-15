package main

import (
	"./animal"
	"./cat"
	"./handler"
	"./interceptor"

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

	// サーバー起動
	e.Start(":8000") //ポート番号指定してね
}
