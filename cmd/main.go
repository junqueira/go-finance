package main

import (
	"time"

	"github.com/FlashBoys/go-finance"
)

func main() {

	// A client for go-finance.

	c, err := finance.NewClient()
	if err != nil {
		panic(err)
	}
	end := time.Now()
	start := end.AddDate(0, -1, 0)

	_, err = c.GetQuoteHistory("YHOO", start, end)
	if err != nil {
		panic(err)
	}

}
