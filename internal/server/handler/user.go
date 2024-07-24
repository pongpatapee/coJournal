package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type UserHTTPHandler struct {
	userService service.UserService
}

func NewUserHTTPHandler(userService service.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.userService.Create(&user)
	if err != nil {
		ErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHTTPHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAll()
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.userService.FindByID(id)
	if err != nil {
		ErrorResponse(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = id

	err = h.userService.Update(&user)
	if err != nil {
		ErrorResponse(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.userService.Delete(id)
	if err != nil {
		ErrorResponse(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
