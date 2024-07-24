package handler

import "net/http"

func ErrorResponse(w http.ResponseWriter, errorString string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	http.Error(w, errorString, statusCode)
}
