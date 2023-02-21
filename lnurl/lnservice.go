package lnurl

import (
	"fmt"

	"github.com/NonsoAmadi10/lightning-web-app/config"
	"github.com/NonsoAmadi10/lightning-web-app/models"
	"github.com/NonsoAmadi10/lightning-web-app/utils"
	LN "github.com/fiatjaf/go-lnurl"
)

func GenerateURL(params LNStruct) string {
	// generate a unique identifier
	identifier := utils.RandomString(5)

	lnurl := fmt.Sprintf("http://localhost:4000/u?q=%v", identifier)

	// Encode the LNURL
	encodedLNURL, err := LN.LNURLEncode(lnurl)
	if err != nil {
		fmt.Println(err)
	}

	newLn := models.LNEntity{
		Identifier:     identifier,
		SatMinSendable: params.MinSendable,
		SatMaxSendable: params.MaxSendable,
		LnurlTag:       "pay",
		Url:            encodedLNURL,
	}

	if err := config.DB.Create(&newLn).Error; err != nil {
		fmt.Printf("An error occured: %v", err)
		panic("Unable to save lnurl")
	}

	return encodedLNURL

}
