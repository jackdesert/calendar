package event

import (
	//"github.com/davecgh/go-spew/spew"
	"log"
	"strconv"
	"strings"
	"time"
)

// See https://golang.org/src/time/format.go
// Note this has to be 2006 for it to work
const dateFormat = "2006-01-02"
const dateFormatPretty = "Monday, Jan 1"

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
type CarouselHolder struct {
	CarouselSlice map[string][]Event
}

func All() []Event {
	return []Event{
		Event{Name: "Jam Skate", Time: "5", Hostess: "Diane", Venue: "The Slots", Address: "513 S Street", DaysOfWeek: "fri", WeeksOfMonth: "1"},
		Event{Name: "hi", Time: "5", Hostess: "Charlie", Venue: "Bourbon", Address: "221 5th Street", DaysOfWeek: "thur", WeeksOfMonth: "5"},
	}
}

func FormattedDate(dateString string) string {
	timeFromDateString, _ := time.Parse(dateFormat, dateString)
	return timeFromDateString.Format(dateFormatPretty)
}

func CarouselInStruct() CarouselHolder {
	return CarouselHolder{
		CarouselSlice: Carousel(),
	}
}

func eventsMatchingDateString(dateString string) []Event {
	events := make([]Event, 0)
	for _, event := range All() {
		log.Println("---")
		log.Println(dateString)
		log.Println(event.Name)
		if event.displayOn(dateString) {
			events = append(events, event)
			log.Println("FOUND")
		}
	}
	return events
}

func Carousel() map[string][]Event {
	dateMap := make(map[string][]Event)
	now := time.Now()
	log.Println("Now()", now)
	for i := 0; i < 2; i++ {
		//log.Println(i)
		//log.Println(now)
		dateString := now.Format(dateFormat)
		//log.Println(dateString)
		dateMap[dateString] = eventsMatchingDateString(dateString)
		now = now.Add(time.Duration(24) * time.Hour)
	}

	return dateMap
}

func (e Event) dayOfWeekMatch(time time.Time) bool {
	weekday := time.Weekday()
	log.Println("weekday: %i", weekday)
	weekdayString := weekday.String()
	log.Println("weekdaystring: %s", weekdayString)
	weekdayAbbreviation := strings.ToLower(weekdayString)[0:3]
	log.Println("weekdayabbreviation: %s", weekdayAbbreviation)
	log.Println("e.DaysOfWeek", e.DaysOfWeek)
	result := strings.Contains(e.DaysOfWeek, weekdayAbbreviation)
	log.Println("result: %s", result)
	return result
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

func (e Event) displayOn(dateString string) bool {

	timeFromDateString, _ := time.Parse(dateFormat, dateString)

	return e.dayOfWeekMatch(timeFromDateString) && e.weekOfMonthMatch(timeFromDateString)
}
