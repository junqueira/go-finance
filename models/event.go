package models

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	// Dividend constant.
	Dividend = "DIVIDEND"
	// Split constant.
	Split = "SPLIT"
)

// Event contains one historical event (either a split or a dividend).
type Event struct {
	Symbol      string
	EventType   string
	Date        time.Time
	DividendAmt decimal.Decimal
	SplitRatio  string
}

// NewEvent creates a new instance of an event struct.
func NewEvent(symbol string, row []string) Event {

	e := Event{Symbol: symbol, EventType: row[0]}

	dateString := parseMalformedDate(row[1])

	// Convert date and set time to market close.
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		fmt.Println("Error serializing historical date: ", err)
		// return e
	}
	e.Date = date.Add(time.Hour * 16)

	if e.EventType == Dividend {
		e.DividendAmt = ToDecimal(row[2])
	} else {
		e.SplitRatio = row[2]
	}

	return e
}
