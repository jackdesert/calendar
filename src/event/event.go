package event

type Event struct {
	Name        string
	Time        string
	Venue       string
	Address     string
	Hostess     string
	DayOfWeek   string
	WeekOfMonth string
	Website     string
}

type EventList struct {
	EventSlice []Event
}
