package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// ToFloat converts an interface string to a float.
func ToFloat(value interface{}) float64 {
	str, ok := value.(string)
	if ok {
		flt, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return flt
		}
	}
	return 0.0
}

// ToInt converts a string to an int.
func ToInt(value string) int {

	if value == "N/A" || value == "-" {
		return 0
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error converting ", value, " to int: ", err)
	}
	return i
}

// ToDecimal converts a string to a decimal value.
func ToDecimal(value string) decimal.Decimal {

	if value == "N/A" || value == "-" {
		return decimal.Decimal{}
	}

	value = strings.Replace(value, "%", "", -1)
	dec, err := decimal.NewFromString(value)
	if err != nil {
		fmt.Println("Error converting ", value, " to decimal value: ", err)
	}
	return dec
}

// ParseDate converts a string to a proper date.
func ParseDate(dString string) time.Time {

	if dString == "N/A" {
		return time.Time{}
	}

	d, err := time.Parse("1/2/2006", dString)
	if err != nil {
		fmt.Println("Error converting ", dString, " to date: ", err)
	}
	return d
}

// ParseDateAndTime converts a string to a proper date with a time.
func ParseDateAndTime(dString string, tString string) time.Time {

	if dString == "N/A" || tString == "N/A" {
		return time.Time{}
	}

	d, err := time.Parse("1/2/2006", dString)
	t, err := time.Parse("3:04pm", tString)
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error converting ", dString, " to date and time: ", err)
	}

	return time.Date(d.Year(), d.Month(), d.Day(), t.Hour(), t.Minute(), t.Second(), 0, loc)
}

func parseMalformedDate(in string) string {

	slice := strings.Split(in, "")
	slice = slice[1:]
	slice = insert(slice, 4, "-")
	slice = insert(slice, 7, "-")

	return strings.Join(slice[:], "")
}

func insert(s []string, i int, x string) []string {

	s = append(s, "")
	copy(s[i+1:], s[i:])
	s[i] = x

	return s
}
