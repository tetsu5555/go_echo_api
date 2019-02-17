package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func MainJwt(c echo.Context) error {
	// contextに定義されているstoreからデータを取得する
	user := c.Get("user")

	// ここどういう処理？
	token := user.(*jwt.Token)

	// ここでmapに変換している
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are on the secret jwt page!")
}
