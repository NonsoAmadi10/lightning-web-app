package main

import (
	"net/http"

	"github.com/NonsoAmadi10/lightning-web-app/lnurl"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSONPretty(http.StatusOK, &Response{Message: "Welcome to lnurl"}, "")
	})

	e.GET("/generate-payment-link", lnurl.GenerateLNURL)
	e.Logger.Fatal(e.Start(":3000"))
}
