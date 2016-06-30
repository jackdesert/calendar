package handler

import (
	//"fmt"
	"github.com/jackdesert/calendar/src/event"
	"html/template"
	"net/http"
)

func events() event.EventList {
	return event.EventList{
		EventSlice: []event.Event{
			event.Event{Name: "Jam Skate", Time: "5", Hostess: "Diane", Venue: "The Slots", Address: "513 S Street", DaysOfWeek: "Friday", WeeksOfMonth: "3"},
			event.Event{Name: "hi", Time: "5", Hostess: "Charlie", Venue: "Bourbon", Address: "221 5th Street", DaysOfWeek: "Tuesday", WeeksOfMonth: "1,3"},
		},
	}
}

func Calendar(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("index.html")
	//defer return500IfError(w)
	t.Execute(w, events())
	//fmt.Fprintf(w, event.render())
}

func return500IfError(w http.ResponseWriter) {
	r := recover() // This returns nil unless there was a panic
	if r != nil {
		http.Error(w, "{\"error\":\"Internal Server Error\"}", 500)
	}
}
