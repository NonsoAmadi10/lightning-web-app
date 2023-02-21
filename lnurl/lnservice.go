package lnurl

import (
	"context"
	"errors"
	"fmt"

	"github.com/NonsoAmadi10/lightning-web-app/config"
	"github.com/NonsoAmadi10/lightning-web-app/models"
	"github.com/NonsoAmadi10/lightning-web-app/utils"
	LN "github.com/fiatjaf/go-lnurl"
	"github.com/lncm/lnd-rpc/v0.10.0/lnrpc"
)

func GenerateURL(params LNStruct) string {
	// generate a unique identifier
	identifier := utils.RandomString(5)

	lnurl := fmt.Sprintf("http://localhost:4000/api/v1/u?q=%v", identifier)

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

func GetIdentifier(identifier string) (interface{}, error) {
	var id models.LNEntity

	var Metadata []string

	// check if the identifier exists
	if err := config.DB.First(&id, "identifier = ?", identifier).RowsAffected; err < 0 {
		return "", errors.New("identifier does not exist")
	}

	if id.LnurlTag == "pay" {
		Metadata = []string{"text/plain", "descriptionTag: pay"}
	} else {
		Metadata = []string{"text/plain", "descriptionTag: Withdraw"}
	}
	// return ln struct

	response := LNStruct{
		Metadata:    Metadata,
		MinSendable: id.SatMinSendable,
		MaxSendable: id.SatMaxSendable,
		Callback:    fmt.Sprintf("http://localhost:4000/api/v1/u/%v", identifier),
		Tag:         id.LnurlTag,
	}

	return response, nil
}

func DecodeLNURL(lnurl string) (string, error) {

	decoded, err := LN.LNURLDecode(lnurl)

	if err != nil {
		return "", errors.New("unable to decode lnurl")
	}

	return decoded, nil
}

func GenPaymentRequest(amount int, identifier string, description string) (string, error) {
	// check if the identifier is a known payment link
	var id models.LNEntity

	if err := config.DB.First(&id, "identifier = ?", identifier).RowsAffected; err < 0 {
		return "", errors.New("identifier does not exist")
	}

	//if it exists, generate a payment invoice

	client := config.Config()

	invoice := &lnrpc.Invoice{
		Memo:   description,
		Value:  int64(amount),
		Expiry: 3600,
	}

	paymentRequest, err := client.AddInvoice(context.Background(), invoice)

	if err != nil {
		return "", errors.New("error generating invoice")
	}

	// save invoice
	newPayReq := models.LNInvoice{
		PaymentID: id.ID,
		Pr:        paymentRequest.PaymentRequest,
	}

	if err := config.DB.Create(&newPayReq).Error; err != nil {
		return "", fmt.Errorf("an error occured creating a payment request %v", err.Error())
	}
	return paymentRequest.PaymentRequest, nil

}
