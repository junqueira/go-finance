package finance

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewExpiration(t *testing.T) {

	// Given we have an date of strings,
	month, day, year := "4", "1", "2016"

	// When we instantiate an expiration,
	e := NewExpiration(month, day, year)

	// And the vars should equal-
	assert.Equal(t, e.Month, "4")
	assert.Equal(t, e.Day, "1")
	assert.Equal(t, e.Year, "2016")
}

func Test_ExpirationFromDate(t *testing.T) {

	// Given we have an date,
	loc, _ := time.LoadLocation("America/New_York")
	date := time.Date(2016, 4, 1, 0, 0, 0, 0, loc)

	// When we instantiate an expiration,
	e := ExpirationFromDate(date)

	// And the vars should equal-
	assert.Equal(t, e.Month, "4")
	assert.Equal(t, e.Day, "1")
	assert.Equal(t, e.Year, "2016")

}
