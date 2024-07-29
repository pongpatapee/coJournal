package handler

import (
	"coJournal/internal/service"

	"github.com/labstack/echo/v4"
)

type NoteHTTPHandler struct {
	noteService service.NoteService
}

func NewNoteHTTPHandler(noteService service.NoteService) *NoteHTTPHandler {
	return &NoteHTTPHandler{
		noteService: noteService,
	}
}

func (h *NoteHTTPHandler) CreateNote(c echo.Context) error {
	return nil
}

func (h *NoteHTTPHandler) GetAllNote(c echo.Context) error {
	return nil
}

func (h *NoteHTTPHandler) GetNote(c echo.Context) error {
	return nil
}

func (h *NoteHTTPHandler) UpdateNote(c echo.Context) error {
	return nil
}

func (h *NoteHTTPHandler) DeleteNote(c echo.Context) error {
	return nil
}
