package ecbank

import (
	"fmt"
	"net/http"

	"github.com/vitostamatti/curconv/money"
)

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrClientSide         = ecbankError("client side error when contacting ECB")
	ErrServerSide         = ecbankError("server side error when contacting ECB")
	ErrUnknownStatusCode  = ecbankError("unknown status code contacting ECB")
	ErrUnexpectedFormat   = ecbankError("unexpected response format")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
)

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

const euroxrefURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type Client struct {
	url string
}

func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {

	if c.url == "" {
		c.url = euroxrefURL
	}

	resp, err := http.Get(c.url)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrCallingServer, err)
	}
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)

	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil
}

func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil

	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)

	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)

	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}

}

func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}
