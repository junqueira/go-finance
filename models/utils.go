package models

import "strconv"

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

// ToInt converts an interface string to an int.
func ToInt(value interface{}) int {
	str, ok := value.(string)
	if ok {
		flt, err := strconv.Atoi(str)
		if err == nil {
			return flt
		}
	}
	return 0.0
}
