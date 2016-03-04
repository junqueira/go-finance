package finance

import (
	"fmt"

	"github.com/FlashBoys/go-finance/models"
)

const quoteURL = "http://download.finance.yahoo.com/d/quotes.csv"

// GetQuote fetches a single symbol's quote from Yahoo Finance.
func GetQuote(symbol string) models.Quote {

	params := map[string]string{
		"s": symbol,
		"f": "saa2bb4c1c4ghjkj1l1m3m4nopp2d1t1vxydee7e8e9j4p5p6qrr1r5r6r7s7t8",
		"e": ".csv",
	}

	table, err := requestCSV(buildURL(quoteURL, params))
	if err != nil {
		fmt.Println("Error fetching quote: ", err)
		return models.Quote{}
	}

	return generateQuotes(table)[0]
}

// GetQuotes fetches multiple symbol's quotes from Yahoo Finance.
func GetQuotes(symbols []string) []models.Quote {

	return []models.Quote{}
}

func generateQuotes(table [][]string) []models.Quote {

	quotes := []models.Quote{}

	for _, row := range table {
		quotes = append(quotes, models.NewQuote(row))
	}

	return quotes

}
