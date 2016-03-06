package models

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

// Bar represents a single bar in time-series of quotes.
type Bar struct {
	Symbol   string
	Date     time.Time
	Open     decimal.Decimal
	High     decimal.Decimal
	Low      decimal.Decimal
	Close    decimal.Decimal
	Volume   int
	AdjClose decimal.Decimal
}

// NewBar creates a new instance of a bar struct.
func NewBar(symbol string, row []string) Bar {

	b := Bar{Symbol: symbol}
	// Convert date and set time to market close.
	date, err := time.Parse("2006-01-02", row[0])
	if err != nil {
		fmt.Println("Error serializing historical date: ", err)
		return b
	}
	date = date.Add(time.Hour * 16)

	b.Date = date
	b.Open = ToDecimal(row[1])
	b.High = ToDecimal(row[2])
	b.Low = ToDecimal(row[3])
	b.Close = ToDecimal(row[4])
	b.Volume = ToInt(row[5])
	b.AdjClose = ToDecimal(row[6])

	return b
}
