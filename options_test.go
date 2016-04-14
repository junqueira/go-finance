package finance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParseExpirations(t *testing.T) {

	// Given we have a slice of timeMaps,
	tms := []timeMap{timeMap{Month: "4", Day: "1", Year: "2016"}}

	// When we parse them to dates,
	dates := parseExpirations(tms)

	// Then the returned slice should have 1 -
	assert.Len(t, dates, 1)

	d := dates[0]

	// And the date should equal-
	assert.Equal(t, 1, d.Day())
	assert.Equal(t, time.April, d.Month())
	assert.Equal(t, 2016, d.Year())

}

func Test_GetOptionsData(t *testing.T) {

	// Given we have some remote options data,
	ts := startTestServer("options_fixture.txt")
	defer ts.Close()

	// When we fetch it,
	result, err := getOptionsData(ts.URL)
	assert.Nil(t, err)

	// Then the price should equal-
	assert.Equal(t, "110.43", result.Price)

	// And the underlying id is-
	assert.Equal(t, "\"22144\"", string(result.Underlying))

	// And the puts/calls/expirations arent empty-
	assert.NotEmpty(t, result.Puts)
	assert.NotEmpty(t, result.Calls)
	assert.NotEmpty(t, result.Expirations)

}
