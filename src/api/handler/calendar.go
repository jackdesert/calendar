package handler

import (
	//"fmt"
	//"github.com/davecgh/go-spew/spew"
	"github.com/jackdesert/calendar/src/event"
	"html/template"
	"log"
	"net/http"
)

func Calendar(w http.ResponseWriter, r *http.Request) {

	carouselInStruct := event.CarouselInStruct()
	//spew.Dump(carousel)

	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"formattedDate": event.FormattedDate,
		"oddOrEven":     event.OddOrEven,
	}
	log.Println(funcMap)

	log.Println("oddoreven", event.OddOrEven)
	event.RestartStripe()

	fileName := "index.html"
	//t, _ := new(template.Template).Funcs(funcMap).ParseFiles(fileName)
	t, _ := template.New("").Funcs(funcMap).ParseFiles(fileName)
	log.Println("Template Defined")
	//t.Funcs(funcMap)
	//log.Println("Funcs Added")
	//defer return500IfError(w)
	t.ExecuteTemplate(w, fileName, carouselInStruct)
	//fmt.Fprintf(w, event.render())
}

func return500IfError(w http.ResponseWriter) {
	r := recover() // This returns nil unless there was a panic
	if r != nil {
		http.Error(w, "{\"error\":\"Internal Server Error\"}", 500)
	}
}
