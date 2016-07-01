package handler

import (
	//"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/jackdesert/calendar/src/event"
	"html/template"
	"net/http"
)

func Calendar(w http.ResponseWriter, r *http.Request) {

	carouselInStruct := event.CarouselInStruct()
	//spew.Dump(carousel)
	t, _ := template.ParseFiles("index.html")
	//defer return500IfError(w)
	t.Execute(w, carouselInStruct)
	//fmt.Fprintf(w, event.render())
}

func return500IfError(w http.ResponseWriter) {
	r := recover() // This returns nil unless there was a panic
	if r != nil {
		http.Error(w, "{\"error\":\"Internal Server Error\"}", 500)
	}
}
