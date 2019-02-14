package main

import (
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

	// ルーティング
	e.GET("/hello/:username", handler.MainPage(), interceptor.BasicAuth())

	e.GET("/cats/:data", cat.GetCat)

	// サーバー起動
	e.Start(":8000") //ポート番号指定してね
}
