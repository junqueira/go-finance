package finance

import (
	"log"
	"time"
)

// Timestamp is a simple time construct.
type Timestamp struct {
	Month  int
	Day    int
	Year   int
	Hour   int
	Minute int
	Second int
}

// NewStamp creates a new instance of Timestamp.
func NewStamp(s string) Timestamp {

	parsedTime, err := time.Parse("1/2/2006", s)
	if err != nil {
		parsedTime, err = time.Parse("3:04pm", s)
		if err != nil {
			parsedTime, err = parseDashedDate(s)
			if err != nil {
				log.Printf("[go-finance] error parsing time: %s", err.Error())
			}
		}
	}

	// Its just a time.
	if parsedTime.Year() == 0 {
		hour, min, sec := parsedTime.Clock()
		return Timestamp{
			Hour:   hour,
			Minute: min,
			Second: sec,
		}
	}

	// Its a day.
	year, month, day := parsedTime.Date()
	return Timestamp{
		Month: int(month),
		Day:   day,
		Year:  year,
	}
}
