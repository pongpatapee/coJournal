package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

		if err.Error() == "users cannot be the same" {
			return echo.NewHTTPError(http.StatusBadRequest, "user_a and user_b cannot be the same")
		}

		return err
	}

	createdJournal, err := h.journalService.FindByID(c.Request().Context(), journal.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createdJournal)
}

func (h *JournalHTTPHandler) GetAllJournal(c echo.Context) error {
	journals, err := h.journalService.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, journals)
}

func (h *JournalHTTPHandler) GetJournal(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	journal, err := h.journalService.FindByID(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("Could not find Journal: %v", id))
	}

	return c.JSON(http.StatusOK, journal)
}

func (h *JournalHTTPHandler) UpdateJournal(c echo.Context) error {
	var journal entities.Journal
	if err := c.Bind(&journal); err != nil {
		return err
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	journal.ID = id

	err = h.journalService.Update(c.Request().Context(), &journal)
	if err != nil {
		if err.Error() == "users cannot be the same" {
			return echo.NewHTTPError(http.StatusBadRequest, "user_a and user_b cannot be the same")
		}
		return err
	}

	updatedJournal, err := h.journalService.FindByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedJournal)
}

func (h *JournalHTTPHandler) DeleteJournal(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	err = h.journalService.Delete(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
