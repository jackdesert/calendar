package event

import (
	"strconv"
	"strings"
	"time"
)

type Event struct {
	Name         string
	Time         string
	Venue        string
	Address      string
	Hostess      string
	DaysOfWeek   string
	WeeksOfMonth string
	Website      string
}

// Used to pass to an html template
type EventList struct {
	EventSlice []Event
}

func (e Event) dayOfWeekMatch(time time.Time) bool {
	weekday := time.Weekday()
	weekdayString := weekday.String()
	weekdayAbbreviation := strings.ToLower(weekdayString)[0:3]
	return strings.Contains(e.DaysOfWeek, weekdayAbbreviation)
}

func (e Event) weekOfMonthMatch(time time.Time) bool {
	if e.WeeksOfMonth == "all" {
		return true
	}

	dayOfMonth := time.Day()
	weekOfMonth := ((dayOfMonth - 1) / 7) + 1
	weekOfMonthString := strconv.Itoa(weekOfMonth)
	return strings.Contains(e.WeeksOfMonth, weekOfMonthString)
}

func (e Event) DisplayOn(dateString string) bool {

	format := "2015-01-01"
	time, _ := time.Parse(format, dateString)

	return e.dayOfWeekMatch(time) && e.weekOfMonthMatch(time)
}
