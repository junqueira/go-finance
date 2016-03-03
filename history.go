package finance

import (
	"fmt"
	"time"

	"github.com/FlashBoys/go-finance/models"
)

// GetQuoteHistory fetches a single security's quote history from Yahoo Finance.
func (c *Client) GetQuoteHistory(symbol string, start time.Time, end time.Time) (bars []models.Bar) {

	// Query YQL for a list of historical prices given input paramaters.
	results, err := c.DB.Query(
		"select * from yahoo.finance.historicaldata where symbol = ? and startDate = ? and endDate = ?",
		symbol,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"))
	if err != nil {
		fmt.Println("Error querying historical: ", err)
		return
	}

	// Serialize results into slice of bars.
	for results.Next() {

		var data map[string]interface{}
		err = results.Scan(&data)
		if err != nil {
			fmt.Println("Error serializing bar: ", err)
			continue
		}

		// Append to slice.
		bars = append(bars, models.NewBar(data))
	}

	return bars
}
