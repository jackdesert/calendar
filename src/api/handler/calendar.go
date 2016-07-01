package handler

import (
	//"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/jackdesert/calendar/src/event"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func Calendar(w http.ResponseWriter, r *http.Request) {

	carouselInStruct := event.CarouselInStruct()
	//spew.Dump(carousel)

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}
	log.Println(funcMap)

	//	t, _ := template.New("abalone").Funcs(funcMap).ParseFiles("index.html")
	t, _ := template.ParseFiles("index.html")
	//t, _ := template.New("abalone").ParseFiles("index.html")
	log.Println("Template Defined")
	//t.Funcs(funcMap)
	//log.Println("Funcs Added")
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
