package finance

import (
	"fmt"
	"strings"

	"github.com/FlashBoys/go-finance/models"
)

const quoteURL = "http://download.finance.yahoo.com/d/quotes.csv"

// GetQuote fetches a single symbol's quote from Yahoo Finance.
func GetQuote(symbol string) models.Quote {

	params := map[string]string{
		"s": symbol,
		"f": strings.Join(models.QuoteFields[:], ""),
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

	params := map[string]string{
		"s": strings.Join(symbols[:], ","),
		"f": strings.Join(models.QuoteFields[:], ""),
		"e": ".csv",
	}

	table, err := requestCSV(buildURL(quoteURL, params))
	if err != nil {
		fmt.Println("Error fetching quote: ", err)
		return []models.Quote{}
	}

	return generateQuotes(table)
}

func generateQuotes(table [][]string) (quotes []models.Quote) {

	for _, row := range table {
		quotes = append(quotes, models.NewQuote(row))
	}
	return quotes
}
