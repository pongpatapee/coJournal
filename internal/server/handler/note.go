package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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
	note := new(entities.Note)

	if err := c.Bind(note); err != nil {
		return err
	}

	if err := c.Validate(note); err != nil {
		return err
	}

	err := h.noteService.Create(c.Request().Context(), note)
	if err != nil {
		return err
	}

	creatednote, err := h.noteService.FindByID(c.Request().Context(), note.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, creatednote)
}

func (h *NoteHTTPHandler) GetAllNote(c echo.Context) error {
	notes, err := h.noteService.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, notes)
}

func (h *NoteHTTPHandler) GetNote(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	note, err := h.noteService.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Could not find Note: %v", id))
	}

	return c.JSON(http.StatusOK, note)
}

func (h *NoteHTTPHandler) UpdateNote(c echo.Context) error {
	return nil
}

func (h *NoteHTTPHandler) DeleteNote(c echo.Context) error {
	return nil
}
