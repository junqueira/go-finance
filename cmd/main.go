package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

	// Fetches the quote for USD/GBP pair.
	quote := finance.GetCurrencyPairQuote(finance.USDGBP)
	fmt.Println(quote)
}
