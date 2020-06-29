package main

import (
	"github.com/rohsa/days-count-go/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.ViewHandler)
	http.HandleFunc("/date", handlers.DateHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
