package animal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Cat struct {
	Name string `json:"name"` //この文字列をタグと呼ばれるもの。reflectパッケージを用いることで変数につけられたタグを取得することができます。
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// echo.Contextは現在のHTTPリクエストのコンテキストを表すもの。
func AddCat(c echo.Context) error {
	cat := new(Cat)

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	// JSON形式の文字列をパースするには、あらかじめJSONのデータ構造に合わせて構造体を定義し、Unmarshal関数を使用する
	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarsharing in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your cat: %#v", cat) //この#は何なんやろ？
	return c.String(http.StatusOK, "we got your cat!")
}

func AddDog(c echo.Context) error {
	dog := new(Dog)

	// どうしてcloseする必要があるのかを調べる
	defer c.Request().Body.Close()

	// ここなにしてるんやろ？
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Failed processing addDog request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your dog: %#v", dog)
	return c.String(http.StatusOK, "we got your dog!")
}

func AddHamster(c echo.Context) error {
	hamster := new(Hamster)

	// 他の二つと比べてパフォーマンスが良くない
	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}
