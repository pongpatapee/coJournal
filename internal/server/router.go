package server

import (
	"coJournal/internal/repository/in_memory"
	"coJournal/internal/server/handler"
	"coJournal/internal/service"
	"fmt"
	"net/http"
)

// TODO: group routes and setup server in server.go, injecting configuration
func SetupRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to CoJournal!")
	})

	// TODO: Refactor with register routes or something

	userRepo := in_memory.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHTTPHandler(userService)

	r.HandleFunc("POST /api/user", userHandler.CreateUser)
	r.HandleFunc("GET /api/user", userHandler.GetAllUser)
	r.HandleFunc("GET /api/user/{id}", userHandler.GetUser)
	r.HandleFunc("PUT /api/user/{id}", userHandler.UpdateUser)
	r.HandleFunc("DELETE /api/user/{id}", userHandler.DeleteUser)

	journalRepo := in_memory.NewInMemoryJournalRepository()
	journalService := service.NewJournalService(journalRepo)
	journalHandler := handler.NewJournalHTTPHandler(journalService)

	r.HandleFunc("POST /api/journal", journalHandler.CreateJournal)
	r.HandleFunc("GET /api/journal", journalHandler.GetAllJournal)
	r.HandleFunc("GET /api/journal/{id}", journalHandler.GetJournal)
	r.HandleFunc("PUT /api/journal/{id}", journalHandler.UpdateJournal)
	r.HandleFunc("DELETE /api/journal/{id}", journalHandler.DeleteJournal)

	noteRepo := in_memory.NewInMemoryNoteRepository()
	noteService := service.NewNoteService(noteRepo)
	noteHandler := handler.NewNoteHTTPHandler(noteService)

	r.HandleFunc("POST /api/note", noteHandler.CreateNote)
	r.HandleFunc("GET /api/note", noteHandler.GetAllNote)
	r.HandleFunc("GET /api/note/{id}", noteHandler.GetNote)
	r.HandleFunc("PUT /api/note/{id}", noteHandler.UpdateNote)
	r.HandleFunc("DELETE /api/note/{id}", noteHandler.DeleteNote)

	return r
}
