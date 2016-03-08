package main

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance"
)

func main() {

	// syms := []string{"AAPL", "TWTR", "FB"}
	// s := finance.GetQuotes(syms)
	// fmt.Println(s)

	// Set time bounds to 1 month starting Jan. 1.
	start, _ := time.Parse(time.RFC3339, "2010-01-01T16:00:00+00:00")
	end := time.Now()

	// Request event history for AAPL.
	e := finance.GetDividendSplitHistory("AAPL", start, end)
	fmt.Println(e)

}
