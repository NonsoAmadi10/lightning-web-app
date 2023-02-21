package lnurl

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NonsoAmadi10/lightning-web-app/config"
	LN "github.com/fiatjaf/go-lnurl"
	"github.com/labstack/echo/v4"
	"github.com/lncm/lnd-rpc/v0.10.0/lnrpc"
)

type LnurlResponse struct {
	Lnurl string `json:"lnurl"`
}
type AddressResponse struct {
	Addr string `json: "addr"`
}

func GenerateLNURL(c echo.Context) error {
	minAmount := c.QueryParam("minAmount")
	maxAmount := c.QueryParam("maxAmount")
	callbackURL := c.QueryParam("callbackURL")
	tag := c.QueryParam("tag")

	//TODO: create constant k1 and save to a db, that constant has to be unique and would be used everytime I need to create a new lnurl

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

func GetAddress(c echo.Context) error {

	client := config.Config()

	addressType := lnrpc.AddressType_NESTED_PUBKEY_HASH

	request := &lnrpc.NewAddressRequest{
		Type: addressType,
	}

	addr, err := client.NewAddress(context.Background(), request)

	if err != nil {
		fmt.Printf("Failed to create new address: %v", err)
		return c.JSONPretty(http.StatusInternalServerError, "", "")
	}

	return c.JSONPretty(http.StatusCreated, &AddressResponse{Addr: addr.Address}, "")
}

// func GetLnAddress(c echo.Context) error {

// }
