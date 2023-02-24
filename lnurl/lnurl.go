package lnurl

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/NonsoAmadi10/lightning-web-app/utils"
	"github.com/labstack/echo/v4"
)

func GenerateLNURL(c echo.Context) error {

	params := LNStruct{
		MinSendable: 10000,
		MaxSendable: 20000,
		Metadata:    []string{"text/plain", "descriptionTag: pay"},
		Tag:         "pay",
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

func GetLNParams(c echo.Context) error {
	// Get Identifier
	identifier := c.QueryParam("q")

	getID, err := GetIdentifier(identifier)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := &utils.SuccessResponse{
		Message: "Payment Parameters retrieved",
		Data:    getID,
	}

	return c.JSONPretty(http.StatusOK, response, "")
}

func Decode(c echo.Context) error {
	// Get url from params
	lnurl := c.QueryParam("url")

	decoded, err := DecodeLNURL(lnurl)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := &utils.SuccessResponse{
		Message: "lnurl decoded successfully",
		Data:    decoded,
	}

	return c.JSONPretty(http.StatusOK, response, "")
}

func GetLNPay(c echo.Context) error {
	var amount int64
	identifier := c.Param("identifier")
	amt, err := strconv.Atoi(c.QueryParam("amount"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "amount must be an integer")
	}
	description := c.QueryParam("desc")

	amount = int64(amt)

	// generate lnpay

	pr, err := GenPaymentRequest(amount, identifier, description)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	successAction := make(map[string]string)

	response := &utils.SuccessResponse{
		Data: &utils.LNPay{
			Pr:            pr,
			Routes:        make([]string, 0),
			SuccessAction: successAction,
		},
	}

	return c.JSONPretty(http.StatusOK, response, "")
}

func GetWURL(c echo.Context) error {
	request, err := GetLNWithdraw()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &utils.SuccessResponse{
		Message: "Here is your withdraw request",
		Data: map[string]string{
			"lnurl": request,
		},
	}

	return c.JSONPretty(http.StatusOK, response, "")
}

func GetWParams(c echo.Context) error {
	identifier := c.Param("identifier")

	amount := 1000

	params, err := GetLNW(amount, identifier)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &utils.SuccessResponse{
		Data: params,
	}

	return c.JSONPretty(http.StatusOK, response, "")
}

func LNWithdrawPay(c echo.Context) error {

	k1 := c.QueryParam("k1")
	payReq := c.QueryParam("pr")

	if k1 == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "k1 is required")
	}

	if payReq == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "pr is required")
	}

	paymentSuccess, err := ProcessLNW(k1, payReq)

	fmt.Println(paymentSuccess)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &utils.SuccessResponse{
		Data:    paymentSuccess,
		Message: "OK",
	}

	return c.JSONPretty(http.StatusOK, response, "")
}

func GetInvoiceByPR(c echo.Context) error {
	pr := c.QueryParam("pr")

	invoice, err := GetInvoice(pr)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := &utils.SuccessResponse{
		Data:    invoice,
		Message: "OK",
	}

	return c.JSONPretty(http.StatusOK, response, "")
}
