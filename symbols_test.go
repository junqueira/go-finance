package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSymbolsFromURL(t *testing.T) {

	// Given that we want to download a list of symbols
	ts := startTestServer("symbols_fixture.csv")
	defer ts.Close()

	// When we request the csv,
	table, err := getSymbolsFromURL(ts.URL)
	assert.Nil(t, err)

	// Then the returned table should have a lot of rows-
	assert.Len(t, table, 647)
}

func Test_ProcessSymbols(t *testing.T) {

	// Given that we have a csv of symbols,
	table := getFixtureAsTable("symbols_fixture.csv")

	// When we parse it,
	symbols := processSymbols(table)

	// Then the returned slice of symbols should contain-
	assert.Contains(t, symbols, "AA")
	assert.Contains(t, symbols, "AAPL")

}
