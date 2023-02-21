package lnurl

import (
	"net/http"

	"github.com/NonsoAmadi10/lightning-web-app/utils"
	"github.com/labstack/echo/v4"
)

func GenerateLNURL(c echo.Context) error {

	params := LNStruct{
		MinSendable: 10000,
		MaxSendable: 20000,
		Metadata:    []string{"text/plain", "descriptionTag: pay"},
	}

	lnurl := GenerateURL(params)

	response := &utils.SuccessResponse{
		Message: "Here is your payment link",
		Data: map[string]string{
			"lnurl": lnurl,
		},
	}

	return c.JSONPretty(http.StatusCreated, response, "")

}

// func GetAddress(c echo.Context) error {

// 	client := config.Config()

// 	addressType := lnrpc.AddressType_NESTED_PUBKEY_HASH

// 	request := &lnrpc.NewAddressRequest{
// 		Type: addressType,
// 	}

// 	addr, err := client.NewAddress(context.Background(), request)

// 	if err != nil {
// 		fmt.Printf("Failed to create new address: %v", err)
// 		return c.JSONPretty(http.StatusInternalServerError, "", "")
// 	}

// 	return c.JSONPretty(http.StatusCreated, &AddressResponse{Addr: addr.Address}, "")
// }
