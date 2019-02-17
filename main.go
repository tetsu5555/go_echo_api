package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"./animal"
	"./cat"
	"./handler"
	"./interceptor"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret admin page!")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the cookie page!")
}

func mainJwt(c echo.Context) error {
	// contextに定義されているstoreからデータを取得する
	user := c.Get("user")

	// ここどういう処理？
	token := user.(*jwt.Token)

	// ここでmapに変換している
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are on the secret jwt page!")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username and password against DB after hashing the password
	if username == "jack" && password == "1234" {
		cookie := &http.Cookie{}

		// this is the same
		//cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		// create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"jack",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("mySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
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

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
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
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "static",
		HTML5: true,
	}))

	// ルーティング
	// 第三引数に渡したミドルウェアでBasicAuthenticationを行なっている
	e.GET("/hello/:username", handler.MainPage(), interceptor.BasicAuth())

	e.GET("/cats/:data", cat.GetCat)

	e.POST("/cats", animal.AddCat)
	e.POST("/dogs", animal.AddDog)
	e.POST("/hamsters", animal.AddHamster)
	e.GET("/login", login)

	// Groupを作成
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	cookieGroup.Use(checkCookie)

	// /jwt にリクエストを送った際にJWTトークンでの認証を行う
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))

	// /admin/maiにリクエストした際に、mainAdminが呼び出されるようになる
	adminGroup.GET("/main", mainAdmin)
	cookieGroup.GET("/main", mainCookie)
	jwtGroup.GET("/main", mainJwt)

	// サーバー起動
	e.Start(":8000") //ポート番号指定してね
}
