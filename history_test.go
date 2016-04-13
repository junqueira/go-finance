package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetHistoryTable(t *testing.T) {

	ts := startTestServer("history_fixture.csv")
	defer ts.Close()

	table, err := getHistoryTable(ts.URL)
	assert.Nil(t, err)

	// Then the returned table should have a lot of rows-
	assert.NotEmpty(t, table)

}

func Test_GenerateBars(t *testing.T) {

	// Given we have a historical csv.
	table := getFixtureAsTable("history_fixture.csv")

	// When we generate historical bars,
	bars := generateBars("AAPL", table)

	// Then the returned slice should have a lot of bar pointers-
	assert.NotEmpty(t, bars)

	first := bars[0]
	second := bars[1]

	// And the first bar symbol should be AAPL.
	assert.Equal(t, "AAPL", first.Symbol)
	// And the second bar symbol should be AAPL.
	assert.Equal(t, "AAPL", second.Symbol)

	// And the length of bars should be the same as rows minus the header-
	assert.Equal(t, len(table)-1, len(bars))
}

func Test_GenerateEvents(t *testing.T) {

	// Given we have a multi-event csv.
	table := getFixtureAsTable("events_fixture.csv")
	// When we generate events,
	events := generateEvents("AAPL", table)
	// Then the returned slice should have a lot of event pointers-
	assert.NotEmpty(t, events)

	first := events[0]
	second := events[1]

	// And the first bar symbol should be AAPL.
	assert.Equal(t, "AAPL", first.Symbol)
	// And the second bar symbol should be AAPL.
	assert.Equal(t, "AAPL", second.Symbol)

	// And the length of events should be the same as rows minus the junk rows-
	assert.Equal(t, 9, len(events))
}
