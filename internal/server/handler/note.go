package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type NoteHTTPHandler struct {
	noteService service.NoteService
}

func NewNoteHTTPHandler(noteService service.NoteService) *NoteHTTPHandler {
	return &NoteHTTPHandler{
		noteService: noteService,
	}
}

func (h *NoteHTTPHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note entities.Note

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.noteService.Create(&note)
	if err != nil {
		ErrorResponse(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *NoteHTTPHandler) GetAllNote(w http.ResponseWriter, r *http.Request) {
	notes, err := h.noteService.FindAll()
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NoteHTTPHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.noteService.FindByID(id)
	if err != nil {
		ErrorResponse(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *NoteHTTPHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	var note entities.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	note.ID = id

	err = h.noteService.Update(&note)
	if err != nil {
		ErrorResponse(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *NoteHTTPHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.noteService.Delete(id)
	if err != nil {
		ErrorResponse(w, "Failed to delete note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
