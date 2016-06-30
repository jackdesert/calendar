package handler

import (
	//"fmt"
	"github.com/jackdesert/calendar/src/event"
	"html/template"
	"net/http"
)

func events() []event.Event {
	return []event.Event{
		event.Event{Name: "Jam Skate", Time: "5", Hostess: "Diane", Venue: "The Slots", Address: "513 S Street", DayOfWeek: "Friday", WeekOfMonth: "3"},
		event.Event{Name: "hi", Time: "5", Hostess: "Charlie", Venue: "Bourbon", Address: "221 5th Street", DayOfWeek: "Tuesday", WeekOfMonth: "1,3"},
	}
}

func Calendar(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("event.html")
	//defer return500IfError(w)
	for _, event := range events() {
		t.Execute(w, event)
		//fmt.Fprintf(w, event.render())
	}
}

func return500IfError(w http.ResponseWriter) {
	r := recover() // This returns nil unless there was a panic
	if r != nil {
		http.Error(w, "{\"error\":\"Internal Server Error\"}", 500)
	}
}
