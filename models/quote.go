package models

import (
	"fmt"
	"reflect"
)

// Quote is the object that is returned for a quote inquiry.
type Quote struct {
	Symbol                      string
	Ask                         string
	AverageDailyVolume          string
	Bid                         string
	BookValue                   string
	Change                      string
	Currency                    string
	DaysLow                     string
	DaysHigh                    string
	YearLow                     string
	YearHigh                    string
	MarketCapitalization        string
	LastTradePriceOnly          string
	FiftydayMovingAverage       string
	TwoHundreddayMovingAverage  string
	Name                        string
	Open                        string
	PreviousClose               string
	ChangeinPercent             string
	LastTradeDate               string
	LastTradeTime               string
	Volume                      string
	StockExchange               string
	DividendYield               string
	DividendShare               string
	EarningsShare               string
	EPSEstimateCurrentYear      string
	EPSEstimateNextYear         string
	EPSEstimateNextQuarter      string
	EBITDA                      string
	PriceSales                  string
	PriceBook                   string
	ExDividendDate              string
	PERatio                     string
	DividendPayDate             string
	PEGRatio                    string
	PriceEPSEstimateCurrentYear string
	PriceEPSEstimateNextYear    string
	ShortRatio                  string
	OneyrTargetPrice            string
}

// NewQuote creates a new instance of a quote.
func NewQuote(row []string) Quote {
	s := Quote{}

	for idx, cell := range row {
		reflect.ValueOf(&s).Elem().Field(idx).SetString(cell)
	}

	fmt.Println(s)
	return s
}
