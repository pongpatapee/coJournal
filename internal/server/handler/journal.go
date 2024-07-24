package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type JournalHTTPHandler struct {
	journalService service.JournalService
}

func NewJournalHTTPHandler(journalService service.JournalService) *JournalHTTPHandler {
	return &JournalHTTPHandler{
		journalService: journalService,
	}
}

func (h *JournalHTTPHandler) CreateJournal(w http.ResponseWriter, r *http.Request) {
	var journal entities.Journal

	err := json.NewDecoder(r.Body).Decode(&journal)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.journalService.Create(&journal)
	if err != nil {
		ErrorResponse(w, "Failed to create journal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *JournalHTTPHandler) GetAllJournal(w http.ResponseWriter, r *http.Request) {
	journals, err := h.journalService.FindAll()
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(journals)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *JournalHTTPHandler) GetJournal(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	journal, err := h.journalService.FindByID(id)
	if err != nil {
		ErrorResponse(w, "Journal not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(journal)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *JournalHTTPHandler) UpdateJournal(w http.ResponseWriter, r *http.Request) {
	var journal entities.Journal
	err := json.NewDecoder(r.Body).Decode(&journal)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	journal.ID = id

	err = h.journalService.Update(&journal)
	if err != nil {
		ErrorResponse(w, "Failed to update journal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *JournalHTTPHandler) DeleteJournal(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.journalService.Delete(id)
	if err != nil {
		ErrorResponse(w, "Failed to delete journal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
