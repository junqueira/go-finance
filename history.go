package finance

import (
	"fmt"
	"strconv"
	"time"
)

const (
	// IntervalDaily daily interval.
	IntervalDaily = "d"
	// IntervalWeekly weekly interval.
	IntervalWeekly = "w"
	// IntervalMonthly monthly interval.
	IntervalMonthly = "m"

	historyURL = "http://ichart.finance.yahoo.com/table.csv"
	divURL     = "http://ichart.finance.yahoo.com/x"
)

// Interval is the duration of the bars returned from the query.
type Interval string

// GetQuoteHistory fetches a single symbol's quote history from Yahoo Finance.
func GetQuoteHistory(symbol string, start time.Time, end time.Time, interval Interval) (bars []Bar) {

	params := map[string]string{
		"s":      symbol,
		"a":      strconv.Itoa(int(start.Month())),
		"b":      strconv.Itoa(start.Day()),
		"c":      strconv.Itoa(start.Year()),
		"d":      strconv.Itoa(int(end.Month())),
		"e":      strconv.Itoa(end.Day()),
		"f":      strconv.Itoa(end.Year()),
		"g":      string(interval),
		"ignore": ".csv",
	}

	table, err := requestCSV(buildURL(historyURL, params))
	if err != nil {
		fmt.Println("Error fetching history: ", err)
		return bars
	}

	return generateBars(symbol, table)
}

func generateBars(symbol string, table [][]string) (bars []Bar) {

	for idx, row := range table {
		if idx != 0 {
			bars = append(bars, NewBar(symbol, row))
		}
	}
	return bars
}

// GetDividendSplitHistory fetches a single symbol's dividend and split history from Yahoo Finance.
func GetDividendSplitHistory(symbol string, start time.Time, end time.Time) (events []Event) {

	params := map[string]string{
		"s":      symbol,
		"a":      strconv.Itoa(int(start.Month())),
		"b":      strconv.Itoa(start.Day()),
		"c":      strconv.Itoa(start.Year()),
		"d":      strconv.Itoa(int(end.Month())),
		"e":      strconv.Itoa(end.Day()),
		"f":      strconv.Itoa(end.Year()),
		"g":      "v",
		"y":      "0",
		"ignore": ".csv",
	}

	table, err := requestCSV(buildURL(divURL, params))
	if err != nil {
		fmt.Println("Error fetching dividend history: ", err)
		return events
	}

	return generateEvents(symbol, table)
}

func generateEvents(symbol string, table [][]string) (events []Event) {
	for _, row := range table {
		if row[0] == Dividend || row[0] == Split {
			events = append(events, NewEvent(symbol, row))
		}
	}
	return events
}
