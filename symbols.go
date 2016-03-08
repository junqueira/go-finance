package finance

import "fmt"

const symbolsURL = "http://www.batstrading.com/market_data/symbol_listing/csv/"

// GetUSEquitySymbols fetches the symbols available through BATS, ~8k symbols.
func GetUSEquitySymbols() []string {

	table, err := requestCSV(symbolsURL)
	if err != nil {
		fmt.Println("Error fetching quote: ", err)
		return []string{}
	}

	return processSymbols(table)
}

func processSymbols(table [][]string) (symbols []string) {

	for idx, row := range table {
		if idx != 0 {
			symbols = append(symbols, row[0])
		}
	}
	return symbols
}
