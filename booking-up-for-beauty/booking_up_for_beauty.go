package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	fixedDate, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}

	return fixedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	layout := "January 2, 2006 15:04:05"

	currentTime, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	if currentTime.Before(now) {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	appointment, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}
	return appointment.Hour() >= 12 && appointment.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."
	layout := "1/2/2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	return parsedDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().UTC().Year()
	layout := "2006-01-02 15:04:05"
	anniversaryDateStr := fmt.Sprintf("%d-09-15 00:00:00", currentYear)
	anniversaryDate, err := time.Parse(layout, anniversaryDateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return anniversaryDate
}
