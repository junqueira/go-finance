package finance

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_ToInt(t *testing.T) {

	// Given that we have a string of a integer.
	intString := "34"
	// When we convert it to a proper int,
	properInt := toInt(intString)
	// Then it should equal its integer value.
	assert.Equal(t, 34, properInt)

	// Given that we have a string of a non-integer.
	notIntString := "-"
	// When we convert it to a proper int,
	zero := toInt(notIntString)
	// Then it be equal to 0.
	assert.Equal(t, 0, zero)

}

func Test_ToDecimal(t *testing.T) {

	// Given that we have a string of a decimal.
	decString := "34.4"
	// When we convert it to a proper decimal,
	properDec := toDecimal(decString)
	// Then it should equal its decimal value.
	assert.Equal(t, decimal.NewFromFloat(34.4), properDec)

	// Given that we have a string of a non-decimal.
	notDecString := "-"
	// When we convert it to a proper decimal,
	zeroDec := toDecimal(notDecString)
	// Then it should equal its decimal value.
	assert.Equal(t, decimal.NewFromFloat(0.0), zeroDec)

	// Given that we have a string of a decimal percent.
	percentString := "0.34%"
	// When we convert it to a proper decimal,
	percent := toDecimal(percentString)
	// Then it should equal its decimal value.
	assert.Equal(t, decimal.NewFromFloat(0.34), percent)

}

func Test_ParseDashedDate(t *testing.T) {

	// Given that we have a string of a date.
	dateString := "2016-04-01"
	// When we convert it to a proper date,
	properDate, err := parseDashedDate(dateString)
	assert.Nil(t, err)
	loc, _ := time.LoadLocation("America/New_York")
	date := time.Date(2016, 4, 1, 0, 0, 0, 0, loc)
	// Then it should equal the date April 1, 2016.
	assert.Equal(t, date.Year(), properDate.Year())
	assert.Equal(t, date.Month(), properDate.Month())
	assert.Equal(t, date.Day(), properDate.Day())

	// Given that we have a string of a non-date.
	nonDateString := "N/A"
	// When we convert it to a proper date,
	badDate, err := parseDashedDate(nonDateString)
	assert.Nil(t, err)

	// Then it should equal the date default.
	assert.Equal(t, time.Time{}, badDate)

}

func Test_ParseMalformedDate(t *testing.T) {

	// Given we have a conjoined date string.
	badString := "020110506"
	// When we convert it to a valid date string,
	validString := parseMalformedDate(badString)
	// Then it should equal a valid date string.
	assert.Equal(t, "2011-05-06", validString)

}
