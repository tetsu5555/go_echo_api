package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	// 静的ファイルのパスを設定
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "static",
		HTML5: true,
	}))

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	// e.Use(middleware.Logger())

	e.Use(middleware.Recover())
	// Groupに対してmiddleware設定する
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Use(serverHeader)

}

// カスタムmiddllewareを作成
// レスポンスヘッダーに書き込みを行うs
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("noHeader", "thisHaveNoMeaning")

		return next(c)
	}
}
