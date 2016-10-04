package format

import (
	"time"
)

// FormatDate is formating a time to a string using an translation of
// month names
// if time is true, the date and time are returned
func Date(date time.Time) string {

	return date.Format("02.") +
		" " +
		date.Format("January") +
		" " +
		date.Format("2006") +
		", " +
		date.Format("15:04")
}
