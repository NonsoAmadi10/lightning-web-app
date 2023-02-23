package lnurl

import (
	"context"
	"errors"
	"fmt"
	"log"

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

func GenPaymentRequest(amount int64, identifier string, description string) (string, error) {
	// check if the identifier is a known payment link
	var id models.LNEntity

	if err := config.DB.First(&id, "identifier = ?", identifier).RowsAffected; err < 0 {
		return "", errors.New("identifier does not exist")
	}

	//check if amount is within the minimum and maximum spendable limit

	//if it exists, generate a payment invoice

	client := config.Config()

	invoice := &lnrpc.Invoice{
		Memo:   description,
		Value:  amount,
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

func GetLNWithdraw() (string, error) {

	id := utils.RandomString(4)
	url := fmt.Sprintf("http://localhost:4000/api/v1/lnwithdraw/%v", id)

	// genrate lnurl

	encoded, err := LN.LNURLEncode(url)

	if err != nil {
		return "", errors.New("error occured generating the lnurl")
	}

	// create ln entity
	ln := models.LNEntity{

		LnurlTag:   "withdrawRequest",
		Url:        url,
		Identifier: id,
	}

	if err := config.DB.Create(&ln).Error; err != nil {
		return "", errors.New("error saving to database")
	}

	return encoded, nil

}

func GetLNW(amount int, identifier string) (*LN.LNURLWithdrawResponse, error) {
	// check if identifier exists
	var id models.LNEntity

	if err := config.DB.First(&id, "identifier = ?", identifier).RowsAffected; err < 0 {
		return &LN.LNURLWithdrawResponse{}, errors.New("identifier does not exist")
	}

	// generate k1
	k1 := utils.RandomString(6)

	request := &LN.LNURLWithdrawResponse{
		Tag:                "withdrawRequest",
		Callback:           "http://localhost:4000/api/v1/withdraw/callback",
		K1:                 k1,
		MaxWithdrawable:    int64(amount),
		MinWithdrawable:    int64(amount),
		DefaultDescription: "A withdraw Request",
	}

	id.K1 = k1
	id.Description = "A Withdraw Request"
	id.SatMaxWithdrawable = int(amount)
	id.SatMinWithdrawable = int(amount)

	// update parameters
	config.DB.Save(&id)
	fmt.Println(id.SatMaxWithdrawable)
	fmt.Println(id.SatMinWithdrawable)

	return request, nil

}

func ProcessLNW(k1 string, pr string) (string, error) {
	// check if k1 exist

	var k models.LNEntity

	result := config.DB.Where("k1 = ?", k1).First(&k)

	// if err := config.DB.First(&k, "k1 = ?", k).RowsAffected; err < 0 {
	// 	return "", errors.New("request not recognized")
	// }

	if result.Error != nil {
		return "", errors.New("request not recognized")
	}

	client := config.Config()
	// decode the pay request
	payReq := &lnrpc.PayReqString{
		PayReq: pr,
	}

	decodeReq, err := client.DecodePayReq(context.Background(), payReq)

	if err != nil {
		return "", errors.New("error decoding request")
	}

	// get amount of satoshi
	amount := decodeReq.NumSatoshis

	// check to see if the amount in the decoded pay request exceeds the satMinWithdrawable

	if int(amount) > k.SatMaxWithdrawable {
		return "", errors.New("payment request exceeds the maximum withdrawable satoshis")
	}

	if int(amount) < k.SatMinWithdrawable {
		return "", errors.New("payment request is lesser the maximum withdrawable satoshis")
	}

	// if all conditions pass, pay the invoice
	payment := &lnrpc.SendRequest{
		PaymentRequest: pr,
	}

	withdrawSuccess, err := client.SendPaymentSync(context.Background(), payment)

	if withdrawSuccess.PaymentError != "" {
		log.Println(err.Error())
		return "", errors.New("expired or paid invoice")
	}

	// save invoice in the database as settled
	paidInvoice := models.LNInvoice{
		PaymentID: k.ID,
		Pr:        pr,
		Status:    "settled",
	}

	if err := config.DB.Create(&paidInvoice).Error; err != nil {
		return "", errors.New("error saving to database")
	}

	return "Payment success", nil
}
