package main

import "github.com/FlashBoys/go-finance"

func main() {

	finance.GetCurrencyPairQuote(finance.USDEUR)
	// start := time.Date(2014, time.February, 1, 16, 1, 0, 0, time.Local)
	// end := time.Now()
	//
	// finance.GetQuoteHistory("AAPL", start, end, finance.IntervalDaily)
	//
	// finance.GetDividendSplitHistory("AAPL", start, end)
}
