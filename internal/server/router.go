package server

import (
	"coJournal/internal/repository"
	"coJournal/internal/server/handler"
	"coJournal/internal/service"
	"fmt"
	"net/http"
)

func NewRouter() *http.ServeMux {
	r := http.NewServeMux()

	userRepo := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHTTPHandler(userService)

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to CoJournal!")
	})

	r.HandleFunc("POST /api/user", userHandler.CreateUser)
	r.HandleFunc("GET /api/user", userHandler.GetAllUser)
	r.HandleFunc("GET /api/user/{id}", userHandler.GetUser)
	r.HandleFunc("PUT /api/user/{id}", userHandler.UpdateUser)
	r.HandleFunc("DELETE /api/user/{id}", userHandler.DeleteUser)

	return r
}
