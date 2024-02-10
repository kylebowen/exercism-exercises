package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/02/2006 15:04:05"

	parsed, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return parsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

	parsed, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return parsed.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"

	parsed, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return 12 <= parsed.Hour() && parsed.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	format := "Monday, January 2, 2006, at 15:04"

	parsed, err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("You have an appointment on %v.", parsed.Format(format))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}