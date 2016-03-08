package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

	// Request all BATS symbols.
	s := finance.GetUSEquitySymbols()
	fmt.Println(s)

}
