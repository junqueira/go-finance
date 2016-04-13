package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewFXPairQuote(t *testing.T) {

	// Given we have a csv with pairs data.
	table := getFixtureAsTable("pair_fixture.csv")

	// When we create a new pairs quote instance,
	pq := newFXPairQuote(table[0])

	// Then pairs quote should have some equal fields-
	assert.Equal(t, USDEUR, pq.Symbol)
	assert.Equal(t, "USD/EUR", pq.PairName)

}
