package finance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_URLParams(t *testing.T) {

	// Given we have an expiration and a symbol,
	e := &Expiration{Month: "4", Day: "1", Year: "2016"}
	symbol := "TWTR"

	// When we convert them to params,
	params := urlParams(symbol, e)

	// Then the returned map should have length-
	assert.Len(t, params, 5)

	// And the entries should equal-
	assert.Equal(t, params["q"], symbol)
	assert.Equal(t, params["expd"], "1")
	assert.Equal(t, params["expm"], "4")
	assert.Equal(t, params["expy"], "2016")
	assert.Equal(t, params["output"], "json")

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
