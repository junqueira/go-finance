package models

import "time"

// Bar represents a single bar in time-series of quotes.
type Bar struct {
	Symbol   string
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int
	AdjClose float64
}

// NewBar creates a new instance of a bar struct.
func NewBar(data map[string]interface{}) Bar {

	// Convert date and set time to market close.
	date, _ := time.Parse("2006-01-02", data["Date"].(string))
	date = date.Add(time.Hour * 16)

	return Bar{data["Symbol"].(string), date, ToFloat(data["Open"]), ToFloat(data["High"]), ToFloat(data["Low"]), ToFloat(data["Close"]), ToInt(data["Volume"]), ToFloat(data["Adj_Close"])}
}
