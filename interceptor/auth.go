package interceptor

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		// usernameがuser、passwordがpasswordであれば認証を許可する
		if username == "user" && password == "password" {
			return true, nil
		}
		return false, nil
	})
}
