package finance

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance/models"
)

// GetQuoteHistory fetches a single security's quote history from Yahoo Finance.
func (c *Client) GetQuoteHistory(symbol string, start time.Time, end time.Time) (bars []models.Bar, err error) {

	startStr := start.Format("2006-01-02")
	endStr := end.Format("2006-01-02")
	fmt.Println(startStr)
	fmt.Println(endStr)

	stmt, err := c.DB.Query(
		"select * from yahoo.finance.historicaldata where symbol = ? and startDate = ? and endDate = ?",
		"YHOO",
		"2009-09-11",
		"2010-03-10")

	if err != nil {
		return nil, err
	}

	bars = []models.Bar{}

	fmt.Println("try")
	for stmt.Next() {
		var data map[string]interface{}
		err = stmt.Scan(&data)

		if err != nil {
			fmt.Println(err)
			continue
		}

		date, _ := time.Parse(data["Date"].(string), "2006-01-02")
		fmt.Println(date)

		bars = append(bars, models.Bar{data["Symbol"].(string), date, data["Open"].(float64), data["High"].(float64), data["Low"].(float64), data["Close"].(float64), data["Volume"].(int), data["Adj_Close"].(float64)})

		// fmt.Printf("%v %v %v %v %v\n", data["Date"], data["Open"], data["High"], data["Low"], data["Close"])
	}

	return bars, nil
}

// map[Symbol:AAPL Date:2009-11-11 Open:204.559996 High:205.000006 Low:201.83 Close:203.250006 Volume:110967500 Adj_Close:26.889665]
