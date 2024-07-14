package router

import (
	"coJournal/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	userHandler := handlers.NewUserHTTPHandler()

	r.HandleFunc("POST /api/user", userHandler.CreateUser)
	r.HandleFunc("GET /api/user", userHandler.GetUsers)
	r.HandleFunc("GET /api/user/{id}", userHandler.GetUser)
	r.HandleFunc("PUT /api/user/{id}", userHandler.UpdateUser)
	r.HandleFunc("DELETE /api/user/{id}", userHandler.DeleteUser)

	// r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Welcome to CoJournal!")
	// })

	return r
}
