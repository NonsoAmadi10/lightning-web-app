package app

import (
	"net/http"

	"github.com/NonsoAmadi10/lightning-web-app/config"
	"github.com/NonsoAmadi10/lightning-web-app/lnurl"
	"github.com/NonsoAmadi10/lightning-web-app/models"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func App() *echo.Echo {

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("/api/v1")

	api.GET("/get-lnurl", lnurl.GenerateLNURL)
	api.GET("/u", lnurl.GetLNParams)
	api.GET("/decoded", lnurl.Decode)
	api.GET("/u/:identifier", lnurl.GetLNPay)

	// Initialize DB
	config.SetupDB(&models.LNEntity{}, &models.LNInvoice{})

	return e

}
