package lnurl

import (
	"fmt"
	"net/http"

	LN "github.com/fiatjaf/go-lnurl"
	"github.com/labstack/echo/v4"
)

type LnurlResponse struct {
	Lnurl string `json:"lnurl"`
}

func GenerateLNURL(c echo.Context) error {
	minAmount := c.QueryParam("minAmount")
	maxAmount := c.QueryParam("maxAmount")
	callbackURL := c.QueryParam("callbackURL")
	tag := c.QueryParam("tag")

	// Validate minAmount, maxAmount and callbackURL here

	// Generate unique LNURL
	lnurl := fmt.Sprintf("http://localhost:3000/lnurlp?minAmount=%s&maxAmount=%s&callbackURL=%s&tag=%s", minAmount, maxAmount, callbackURL, tag)

	// Encode the LNURL
	encodedLNURL, err := LN.LNURLEncode(lnurl)
	if err != nil {
		fmt.Println(err)
	}

	return c.JSONPretty(http.StatusCreated, &LnurlResponse{Lnurl: encodedLNURL}, "")
}
