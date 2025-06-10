package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"

	"github.com/vitostamatti/curconv/money"
)

const baseCurrencyCode = "EUR"

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float64 `xml:"rate,attr"`
}

func (e envelope) exchangeRates() map[string]float64 {
	rates := make(map[string]float64, len(e.Rates)+1)
	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}
	rates[baseCurrencyCode] = 1.
	return rates
}

func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		one, err := money.ParseDecimal("1")
		if err != nil {
			return money.ExchangeRate{}, fmt.Errorf("unable to create a rate of value 1: %w", err)
		}
		return money.ExchangeRate(one), nil
	}
	rates := e.exchangeRates()

	sorceFactor, ok := rates[source]
	if !ok {
		return money.ExchangeRate{}, fmt.Errorf("unknown currency %s", source)
	}

	targetFactor, ok := rates[target]
	if !ok {
		return money.ExchangeRate{}, fmt.Errorf("unknown currency %s", target)
	}

	rate, err := money.ParseDecimal(fmt.Sprintf("%.10f", targetFactor/sorceFactor))
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("unable to create a rate of value %f: %w", targetFactor/sorceFactor, err)
	}
	return money.ExchangeRate(rate), nil
}

func readRateFromResponse(source, target string, body io.Reader) (money.ExchangeRate, error) {
	decoder := xml.NewDecoder(body)
	var ratesEnvelope envelope
	err := decoder.Decode(&ratesEnvelope)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}
	rate, err := ratesEnvelope.exchangeRate(source, target)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}
	return rate, nil
}
