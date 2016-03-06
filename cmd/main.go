package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

	syms := []string{"AAPL", "TWTR", "FB"}
	s := finance.GetQuotes(syms)
	fmt.Println(s)

	// // Set time bounds to 1 month starting Jan. 1.
	// start, _ := time.Parse(time.RFC3339, "2016-01-01T16:00:00+00:00")
	// end := start.AddDate(0, 1, 0)
	//
	// // Request daily history for TWTR.
	// // IntervalDaily OR IntervalWeekly OR IntervalMonthly are supported.
	// bars := finance.GetQuoteHistory("TWTR", start, end, finance.IntervalDaily)
	// fmt.Println(bars)

}
