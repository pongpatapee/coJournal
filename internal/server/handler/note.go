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
	var notes []*entities.Note
	var err error

	if c.QueryParam("journal_id") == "" {
		notes, err = h.noteService.FindAll(c.Request().Context())
		if err != nil {
			return err
		}
	} else {
		journalID, err := uuid.Parse(c.QueryParam("journal_id"))
		if err != nil {
			return err
		}

		notes, err = h.noteService.FindByJournalID(c.Request().Context(), journalID)
		if err != nil {
			return err
		}
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
	var note entities.Note

	if err := c.Bind(&note); err != nil {
		return err
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	note.ID = id

	if err := h.noteService.Update(c.Request().Context(), &note); err != nil {
		return c.String(http.StatusNotFound, "Could not find note to update")
	}

	updatednote, err := h.noteService.FindByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatednote)
}

func (h *NoteHTTPHandler) DeleteNote(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	err = h.noteService.Delete(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
