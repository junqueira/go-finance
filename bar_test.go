package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewBar(t *testing.T) {

	// Given we have a csv with historical data.
	table := getFixtureAsTable("history_fixture.csv")

	// When we create a new historical chart bar,
	bar := newBar("AAPL", table[17])

	// Then bar should have some equal fields-
	assert.Equal(t, "AAPL", bar.Symbol)
	assert.Equal(t, 43402300, bar.Volume)

}
