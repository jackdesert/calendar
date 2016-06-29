package main

import (
	"github.com/jackdesert/calendar/src/api/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.Calendar)

	log.Println("Listening1...")
	http.ListenAndServe(":3100", nil)

}
