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

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret admin page!")
}

// カスタムmiddllewareを作成
// レスポンスヘッダーに書き込みを行うs
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("noHeader", "thisHaveNoMeaning")

		return next(c)
	}
}

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	e.Use(ServerHeader)

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Groupに対してmiddleware設定する
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// 静的ファイルのパスを設定
	e.Static("/static", "static")

	// ルーティング
	// 第三引数に渡したミドルウェアでBasicAuthenticationを行なっている
	e.GET("/hello/:username", handler.MainPage(), interceptor.BasicAuth())

	e.GET("/cats/:data", cat.GetCat)

	e.POST("/cats", animal.AddCat)
	e.POST("/dogs", animal.AddDog)
	e.POST("/hamsters", animal.AddHamster)

	// Groupを作成
	g := e.Group("/admin")

	// /admin/maiにリクエストした際に、mainAdminが呼び出されるようになる
	g.GET("/main", mainAdmin)

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
