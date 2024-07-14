package main

import (
	"coJournal/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
