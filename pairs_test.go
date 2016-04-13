package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetPairsQuoteTable(t *testing.T) {

	ts := startTestServer("pair_fixture.csv")
	defer ts.Close()

	table, err := getPairsQuotesTable(ts.URL)
	assert.Nil(t, err)

	// Then the returned slice should have a length of 1.
	assert.Len(t, table, 1)

	sym := table[0][0]

	// And the first row's pair should be USDEUR.
	assert.Equal(t, USDEUR, sym)

}

func Test_GeneratePairQuotes(t *testing.T) {

	// Given we have a pair quote csv.
	table := getFixtureAsTable("pair_fixture.csv")

	// When we generate quotes,
	pairs := generatePairQuotes(table)

	// Then the returned slice of pair quote pointers should have a length of 1.
	assert.Len(t, pairs, 1)

	usdeur := pairs[0]

	// And the first quote symbol should be AAPL.
	assert.Equal(t, USDEUR, usdeur.Symbol)

}
