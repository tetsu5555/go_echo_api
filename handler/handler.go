package handler

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type Animal struct {
	Name string
	Age  int
}

func MainPage() echo.HandlerFunc {
	cat := Animal{"kojirou", 1}
	cat2, err := json.Marshal(cat)
	if err != nil {
		panic(err)
	}
	return func(c echo.Context) error {
		return c.String(http.StatusOK, string(cat2))
	}
}
