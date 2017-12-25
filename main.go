package main

import (
	"fmt"
	"./bittrex"
)

var (
	currencies []bittrex.Currency
	markets []bittrex.Market
	err error
)

func main() {
	api := bittrex.New()

	if currencies, err = bittrex.GetCurrencies(api.Client); err != nil {
		fmt.Println(err)
	}
	fmt.Println(currencies)

	if markets, err = bittrex.GetMarkets(api.Client); err != nil {
		fmt.Println(err)
	}
}
