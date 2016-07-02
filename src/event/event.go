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
const dateFormatPretty = "Monday, Jan 2"

var stripeCounter = 0

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
		Event{Name: "Jam Skate",
			Time:         "8pm",
			Hostess:      "Diane",
			Venue:        "Skadium",
			Address:      "1311 S Bowman Rd, Little Rock, AR",
			DaysOfWeek:   "sun",
			WeeksOfMonth: "all"},
		Event{Name: "Farmer's Market",
			Time:         "8am - noon",
			Hostess:      "",
			Venue:        "North Little Rock",
			Address:      "",
			DaysOfWeek:   "sat",
			WeeksOfMonth: "all"},
		Event{Name: "CoDa",
			Time:         "6:30pm",
			Hostess:      "Miriam",
			Venue:        "Central Church of Christ",
			Address:      "823 W 6th St, Little Rock, AR",
			DaysOfWeek:   "tues",
			WeeksOfMonth: "all"},
		Event{Name: "Open Mic (House of Art)",
			Time:         "9pm",
			Hostess:      "Chris James",
			Venue:        "House of Art",
			Address:      "North Little Rock",
			DaysOfWeek:   "fri",
			WeeksOfMonth: "all"},
		Event{Name: "Free Hair Cuts (House of Art)",
			Time:         "?? 10am - 12pm ??",
			Hostess:      "??",
			Venue:        "House of Art",
			Address:      "North Little Rock",
			DaysOfWeek:   "sat",
			WeeksOfMonth: "3"},
		Event{Name: "Art Walk (North Little Rock)",
			Time:         "5pm - 8pm",
			Hostess:      "",
			Venue:        "Various, including House of Art",
			Address:      "North Little Rock",
			DaysOfWeek:   "fri",
			WeeksOfMonth: "3"},
		Event{Name: "Art Walk",
			Time:         "5pm - 8pm",
			Hostess:      "",
			Venue:        "River Market District plus Free Art Shuttle",
			Address:      "Little Rock",
			DaysOfWeek:   "fri",
			WeeksOfMonth: "2"},
		Event{Name: "Wine & Cheese",
			Time:         "4pm - close",
			Hostess:      "",
			Venue:        "Crush Wine Bar",
			Address:      "North Little Rock",
			DaysOfWeek:   "tues, wed, thurs, fri, sat",
			WeeksOfMonth: "all"},
		Event{Name: "Club Level",
			Time:         "8pm - 2am",
			Hostess:      "",
			Venue:        "Club Level",
			Address:      "315 Main St, Little Rock, AR",
			DaysOfWeek:   "fri, sat",
			WeeksOfMonth: "all"},
		Event{Name: "Standup Comedy",
			Time:         "??--call first",
			Hostess:      "",
			Venue:        "The Joint",
			Address:      "301 Main St, North Little Rock, AR",
			DaysOfWeek:   "tues",
			WeeksOfMonth: "all"},
		Event{Name: "Comedy Improv",
			Time:         "??--call first",
			Hostess:      "",
			Venue:        "The Joint",
			Address:      "301 Main St, North Little Rock, AR",
			DaysOfWeek:   "wed",
			WeeksOfMonth: "all"},
		Event{Name: "Artist Series (The Joint)",
			Time:         "??",
			Hostess:      "",
			Venue:        "The Joint",
			Address:      "301 Main St, North Little Rock, AR",
			DaysOfWeek:   "thurs",
			WeeksOfMonth: "all"},
		Event{Name: "Music & Comedy",
			Time:         "??",
			Hostess:      "",
			Venue:        "The Joint",
			Address:      "301 Main St, North Little Rock, AR",
			DaysOfWeek:   "thurs",
			WeeksOfMonth: "4"},
	}
}

func FormattedDate(dateString string) string {
	timeFromDateString, _ := time.Parse(dateFormat, dateString)
	return timeFromDateString.Format(dateFormatPretty)
}

func RestartStripe() {
	stripeCounter = 1
}

func OddOrEven() string {
	// TODO This will stripe fine if only one person accesses server at a time ;)
	stripeCounter += 1
	if (stripeCounter % 2) == 1 {
		return "odd"
	} else {
		return "even"
	}
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
	for i := 0; i < 14; i++ {
		//log.Println(i)
		//log.Println(now)
		dateString := now.Format(dateFormat)
		//log.Println(dateString)
		dateMap[dateString] = eventsMatchingDateString(dateString)
		now = now.Add(time.Duration(24) * time.Hour)
	}

	return dateMap
}

func (e Event) Frequency() string {
	numberMap := map[string]string{
		"1": "First",
		"2": "Second",
		"3": "Third",
		"4": "Fourth",
		"5": "Fifth",
	}
	output := ""

	numberSlice := strings.Split(e.WeeksOfMonth, ",")

	if e.WeeksOfMonth == "all" {
		output = "Every "

	} else {

		for _, number := range numberSlice {
			output += numberMap[number] + " &"
		}
	}
	// Remove trailing "&"
	if strings.Contains(output, "&") {
		output = output[0 : len(output)-1]
	}

	output += e.DaysOfWeek
	return output
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

func (e Event) weekOfMonthMatch(t time.Time) bool {
	if e.WeeksOfMonth == "all" {
		return true
	}

	log.Println("formatted Date: ", t.Format(dateFormat))
	dayOfMonth := t.Day()
	log.Println("day of month: ", dayOfMonth)
	weekOfMonth := ((dayOfMonth - 1) / 7) + 1
	log.Println("weekOfMonth: ", weekOfMonth)
	weekOfMonthString := strconv.Itoa(weekOfMonth)
	return strings.Contains(e.WeeksOfMonth, weekOfMonthString)
}

func (e Event) displayOn(dateString string) bool {

	timeFromDateString, _ := time.Parse(dateFormat, dateString)

	return e.dayOfWeekMatch(timeFromDateString) && e.weekOfMonthMatch(timeFromDateString)
}
