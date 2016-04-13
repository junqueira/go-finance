package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEvent(t *testing.T) {

	// Given we have a csv with events data.
	table := getFixtureAsTable("events_fixture.csv")

	// When we create a new event,
	event := newEvent("AAPL", table[5])

	// Then bar should have some equal fields-
	assert.Equal(t, "AAPL", event.Symbol)
	assert.Equal(t, Dividend, event.EventType)

}
