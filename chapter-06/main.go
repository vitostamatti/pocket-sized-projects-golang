package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vitostamatti/curconv/ecbank"
	"github.com/vitostamatti/curconv/money"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "", "target currency")
	flag.Parse()

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s", *from, err)
		flag.Usage()
		os.Exit(1)
	}
	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse target currency %q: %s", *to, err)
		flag.Usage()
		os.Exit(1)
	}

	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s\n", value, err)
		flag.Usage()
		os.Exit(1)
	}

	amount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to create amount: %s\n", err)
		os.Exit(1)
	}

	rates := ecbank.Client{}

	convertedAmount, err := money.Convert(amount, toCurrency, rates)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to convert amount: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s = %s\n", amount, convertedAmount)
}
