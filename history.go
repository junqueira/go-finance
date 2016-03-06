package finance

import (
	"fmt"
	"strconv"
	"time"

	"github.com/FlashBoys/go-finance/models"
)

const (
	// IntervalDaily daily interval.
	IntervalDaily = "d"
	// IntervalWeekly weekly interval.
	IntervalWeekly = "w"
	// IntervalMonthly monthly interval.
	IntervalMonthly = "m"

	historyURL = "http://ichart.finance.yahoo.com/table.csv"
)

// Interval is the duration of the bars returned from the query.
type Interval string

// GetQuoteHistory fetches a single Quote's quote history from Yahoo Finance.
func GetQuoteHistory(symbol string, start time.Time, end time.Time, interval Interval) (bars []models.Bar) {

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

func generateBars(symbol string, table [][]string) (bars []models.Bar) {

	for idx, row := range table {
		//fmt.Println(row)
		if idx != 0 {
			bars = append(bars, models.NewBar(symbol, row))
		}
	}
	return bars
}
