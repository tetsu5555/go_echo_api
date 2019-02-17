package main

import (
	"./router"
)

func main() {
	// サーバー起動
	e := router.New()

	e.Start(":8000") //ポート番号指定してね
}
