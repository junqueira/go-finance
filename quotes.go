package finance

import "strings"

const quoteURL = "http://download.finance.yahoo.com/d/quotes.csv"

// GetQuote fetches a single symbol's quote from Yahoo Finance.
func GetQuote(symbol string) (*Quote, error) {

	params := map[string]string{
		"s": symbol,
		"f": strings.Join(quoteFields[:], ""),
		"e": ".csv",
	}

	table, err := requestTable(quoteURL, params)
	if err != nil {
		return nil, err
	}
	return generateQuotes(table)[0], nil
}

// GetQuotes fetches multiple symbol's quotes from Yahoo Finance.
func GetQuotes(symbols []string) ([]*Quote, error) {

	params := map[string]string{
		"s": strings.Join(symbols[:], ","),
		"f": strings.Join(quoteFields[:], ""),
		"e": ".csv",
	}

	table, err := requestTable(quoteURL, params)
	if err != nil {
		return nil, err
	}
	return generateQuotes(table), nil
}

// generateQuotes turns the raw table data of quotes into proper quote structs.
func generateQuotes(table [][]string) (quotes []*Quote) {

	for _, row := range table {
		quotes = append(quotes, NewQuote(row))
	}
	return quotes
}
