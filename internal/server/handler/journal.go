package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type JournalHTTPHandler struct {
	journalService service.JournalService
}

func NewJournalHTTPHandler(journalService service.JournalService) *JournalHTTPHandler {
	return &JournalHTTPHandler{
		journalService: journalService,
	}
}

func (h *JournalHTTPHandler) CreateJournal(c echo.Context) error {
	journal := new(entities.Journal)

	if err := c.Bind(journal); err != nil {
		return err
	}

	if err := c.Validate(journal); err != nil {
		return err
	}

	err := h.journalService.Create(c.Request().Context(), journal)
	if err != nil {
		return err
	}

	createdJournal, err := h.journalService.FindByID(c.Request().Context(), journal.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createdJournal)
}

func (h *JournalHTTPHandler) GetAllJournal(c echo.Context) error {
	return nil
}

func (h *JournalHTTPHandler) GetJournal(c echo.Context) error {
	return nil
}

func (h *JournalHTTPHandler) UpdateJournal(c echo.Context) error {
	return nil
}

func (h *JournalHTTPHandler) DeleteJournal(c echo.Context) error {
	return nil
}
