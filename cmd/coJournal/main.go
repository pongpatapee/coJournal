package main

import (
	"coJournal/internal/server"
	"log"
	"net/http"
	// _ "github.com/lib/pq"
)

func main() {
	r := server.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
