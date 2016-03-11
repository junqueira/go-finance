package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

	// Request all BATS symbols.
	q := finance.GetCurrencyPairQuote(finance.USDGBP)
	fmt.Println(q)

}
