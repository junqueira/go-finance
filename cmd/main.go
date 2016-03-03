package main

import (
	"fmt"

	"github.com/FlashBoys/go-finance"
)

func main() {

	// Init a client.
	c, err := finance.NewClient()
	if err != nil {
		panic(err)
	}

	// Request history for TWTR.
	s := c.GetQuote("TWTR")
	fmt.Println(s)

}
