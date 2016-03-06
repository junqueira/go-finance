package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// QuoteFields are the requested quote fields.
var QuoteFields = []string{
	"s",  // Symbol
	"a",  // Ask
	"a2", // AverageDailyVolume
	"b",  // Bid
	"b4", // BookValue
	"c1", // Change
	"c4", // Currency
	"g",  // DaysLow
	"h",  // DaysHigh
	"j",  // YearLow
	"k",  // YearHigh
	"j1", // MarketCapitalization
	"l1", // LastTradePriceOnly
	"m3", // FiftydayMovingAverage
	"m4", // TwoHundreddayMovingAverage
	"n",  // Name
	"o",  // Open
	"p",  // PreviousClose
	"p2", // ChangeinPercent
	"d1", // LastTradeDate
	"t1", // LastTradeTime
	"v",  // Volume
	"x",  // StockExchange
	"y",  // DividendYield
	"d",  // DividendShare
	"e",  // EarningsShare
	"e7", // EPSEstimateCurrentYear
	"e8", // EPSEstimateNextYear
	"e9", // EPSEstimateNextQuarter
	"j4", // EBITDA
	"p5", // PriceSales
	"p6", // PriceBook
	"q",  // ExDividendDate
	"r",  // PERatio
	"r1", // DividendPayDate
	"r5", // PEGRatio
	"s7", // ShortRatio
	"t8", // OneyrTargetPrice
}

// Quote is the object that is returned for a quote inquiry.
type Quote struct {
	Symbol             string
	Name               string
	LastTradeTime      time.Time
	LastTradePrice     decimal.Decimal
	Ask                decimal.Decimal
	Bid                decimal.Decimal
	Volume             int
	ChangeNominal      decimal.Decimal
	ChangePercent      decimal.Decimal
	Open               decimal.Decimal
	PreviousClose      decimal.Decimal
	Exchange           string
	DayLow             decimal.Decimal
	DayHigh            decimal.Decimal
	FiftyTwoWeekLow    decimal.Decimal
	FiftyTwoWeekHigh   decimal.Decimal
	Currency           string
	MarketCap          string
	FiftyDayMA         decimal.Decimal
	TwoHundredDayMA    decimal.Decimal
	AvgDailyVolume     int
	FiftyTwoWeekTarget decimal.Decimal
	ShortRatio         decimal.Decimal
	BookValue          decimal.Decimal
	EBITDA             string
	PriceSales         decimal.Decimal
	PriceBook          decimal.Decimal
	PERatio            decimal.Decimal
	PEGRatio           decimal.Decimal
	DivYield           decimal.Decimal
	DivPerShare        decimal.Decimal
	DivExDate          time.Time
	DivPayDate         time.Time
	EPS                decimal.Decimal
	EPSEstCurrentYear  decimal.Decimal
	EPSEstNextYear     decimal.Decimal
	EPSEstNextQuarter  decimal.Decimal
}

// NewQuote creates a new instance of a quote.
func NewQuote(row []string) Quote {

	fields := make(map[string]string, 0)
	for idx, v := range QuoteFields {
		fields[v] = row[idx]
	}

	return Quote{
		Symbol:             fields["s"],
		Name:               fields["n"],
		LastTradeTime:      ParseDateAndTime(fields["d1"], fields["t1"]),
		LastTradePrice:     ToDecimal(fields["l1"]),
		Ask:                ToDecimal(fields["a"]),
		Bid:                ToDecimal(fields["b"]),
		Volume:             ToInt(fields["v"]),
		ChangeNominal:      ToDecimal(fields["c1"]),
		ChangePercent:      ToDecimal(fields["p2"]),
		Open:               ToDecimal(fields["o"]),
		PreviousClose:      ToDecimal(fields["p"]),
		Exchange:           fields["x"],
		DayLow:             ToDecimal(fields["g"]),
		DayHigh:            ToDecimal(fields["h"]),
		FiftyTwoWeekLow:    ToDecimal(fields["j"]),
		FiftyTwoWeekHigh:   ToDecimal(fields["k"]),
		Currency:           fields["c4"],
		MarketCap:          fields["j1"],
		FiftyDayMA:         ToDecimal(fields["m3"]),
		TwoHundredDayMA:    ToDecimal(fields["m4"]),
		AvgDailyVolume:     ToInt(fields["a2"]),
		FiftyTwoWeekTarget: ToDecimal(fields["t8"]),
		ShortRatio:         ToDecimal(fields["s7"]),
		BookValue:          ToDecimal(fields["b4"]),
		EBITDA:             fields["j4"],
		PriceSales:         ToDecimal(fields["p5"]),
		PriceBook:          ToDecimal(fields["p6"]),
		PERatio:            ToDecimal(fields["r"]),
		PEGRatio:           ToDecimal(fields["r5"]),
		DivYield:           ToDecimal(fields["y"]),
		DivPerShare:        ToDecimal(fields["d"]),
		DivExDate:          ParseDate(fields["q"]),
		DivPayDate:         ParseDate(fields["r1"]),
		EPS:                ToDecimal(fields["e"]),
		EPSEstCurrentYear:  ToDecimal(fields["e7"]),
		EPSEstNextYear:     ToDecimal(fields["e8"]),
		EPSEstNextQuarter:  ToDecimal(fields["e9"]),
	}
}
