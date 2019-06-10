package main

import (
	"api/router"
)

func main() {
	// サーバー起動
	e := router.New()

	e.Start(":3000") //ポート番号指定してね
}
