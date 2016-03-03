package finance

import (
	"fmt"

	"github.com/FlashBoys/go-finance/models"
)

// GetQuote fetches a single security's quote from Yahoo Finance.
func (c *Client) GetQuote(symbol string) (s models.Security) {

	// Query YQL for a list of historical prices given input paramaters.
	results, err := c.DB.Query(
		"select * from yahoo.finance.quotes where symbol in (\"AAPL\")")
	if err != nil {
		fmt.Println("Error querying quote: ", err)
		return
	}
	fmt.Println(results)
	//
	// // Serialize results into slice of bars.
	// for results.Next() {
	//
	// 	var data map[string]interface{}
	// 	err = results.Scan(&data)
	// 	if err != nil {
	// 		fmt.Println("Error serializing bar: ", err)
	// 		continue
	// 	}
	//
	// }

	return s
}

// GetQuotes fetches multiples security's quotes from Yahoo Finance.
func (c *Client) GetQuotes(symbols []string) {

}
