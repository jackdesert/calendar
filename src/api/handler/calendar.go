package handler

import (
	//"github.com/jackdesert/calendar/src/event"
	"fmt"
	"net/http"
)

func Calendar(w http.ResponseWriter, r *http.Request) {

	defer return500IfError(w)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func return500IfError(w http.ResponseWriter) {
	r := recover() // This returns nil unless there was a panic
	if r != nil {
		http.Error(w, "{\"error\":\"Internal Server Error\"}", 500)
	}
}
