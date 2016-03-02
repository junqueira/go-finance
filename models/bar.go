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
