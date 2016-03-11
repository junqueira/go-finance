package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// FXFields are the requested fx pair fields.
var FXFields = []string{
	"s",  // Symbol
	"n",  // Name
	"d1", // LastTradeDate
	"t1", // LastTradeTime
	"l1", // LastTradePriceOnly
	"c1", // Change
	"p2", // ChangeinPercent
	"g",  // DaysLow
	"h",  // DaysHigh
	"j",  // YearLow
	"k",  // YearHigh
}

// FXPairQuote represents the quote of a currency pair.
type FXPairQuote struct {
	Symbol           string
	PairName         string
	LastTime         time.Time
	LastRate         decimal.Decimal
	ChangeNominal    decimal.Decimal
	ChangePercent    decimal.Decimal
	DayLow           decimal.Decimal
	DayHigh          decimal.Decimal
	FiftyTwoWeekLow  decimal.Decimal
	FiftyTwoWeekHigh decimal.Decimal
}

// NewFXPairQuote creates a new instance of quote of a currency pair.
func NewFXPairQuote(row []string) FXPairQuote {

	fields := make(map[string]string, 0)
	for idx, v := range QuoteFields {
		fields[v] = row[idx]
	}

	return FXPairQuote{
		Symbol:           fields["s"],
		PairName:         fields["n"],
		LastTime:         ParseDateAndTime(fields["d1"], fields["t1"]),
		LastRate:         ToDecimal(fields["l1"]),
		ChangeNominal:    ToDecimal(fields["c1"]),
		ChangePercent:    ToDecimal(fields["p2"]),
		DayLow:           ToDecimal(fields["g"]),
		DayHigh:          ToDecimal(fields["h"]),
		FiftyTwoWeekLow:  ToDecimal(fields["j"]),
		FiftyTwoWeekHigh: ToDecimal(fields["k"]),
	}
}
