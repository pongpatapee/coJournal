package main

import (
	"coJournal/internal/server"
	"log"
	"net/http"
)

func main() {
	r := server.SetupRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
