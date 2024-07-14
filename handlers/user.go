package handlers

import (
	"coJournal/services"
	"encoding/json"
	"net/http"
)

type UserHTTPHandler struct {
	userService services.UserService
}

func NewUserHTTPHandler() *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: *services.NewUserService(),
	}
}

func (handler *UserHTTPHandler) errorResponse(w http.ResponseWriter, statusCode int, errorString string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	// TODO: Add 500 internal server error
	http.Error(w, errorString, statusCode)
}

func (handler *UserHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user services.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handler.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	newUser := handler.userService.Create(user)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		handler.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (handler *UserHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := handler.userService.GetAll()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		handler.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (handler *UserHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := services.UserID(r.PathValue("id"))

	user, found := handler.userService.Get(id)
	if !found {
		handler.errorResponse(w, http.StatusNotFound, "Not Found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		handler.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (handler *UserHTTPHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := services.UserID(r.PathValue("id"))

	var newUser services.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		handler.errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, found := handler.userService.Update(id, newUser)
	if !found {
		handler.errorResponse(w, http.StatusNotFound, "Not Found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		handler.errorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (handler *UserHTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := services.UserID(r.PathValue("id"))

	found := handler.userService.Delete(id)
	if !found {
		handler.errorResponse(w, http.StatusNotFound, "Not Found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
