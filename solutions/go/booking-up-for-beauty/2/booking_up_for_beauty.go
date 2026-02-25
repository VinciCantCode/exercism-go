package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("parse error:", err)
		return false
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	formatedTime := t.Format("Monday, January 2, 2006, at 15:04")
	return "You have an appointment on " + formatedTime + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
	return t
}
