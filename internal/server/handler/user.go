package handler

import (
	"coJournal/internal/entities"
	"coJournal/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHTTPHandler struct {
	userService service.UserService
}

func NewUserHTTPHandler(userService service.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (h *UserHTTPHandler) CreateUser(c echo.Context) error {
	var user entities.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := h.userService.Create(&user); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, &user)
}

func (h *UserHTTPHandler) GetAllUser(c echo.Context) error {
	users, err := h.userService.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHTTPHandler) GetUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := h.userService.FindByID(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Could not find user")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHTTPHandler) UpdateUser(c echo.Context) error {
	var user entities.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	user.ID = id

	if err := h.userService.Update(&user); err != nil {
		return c.String(http.StatusNotFound, "Could not find user to update")
	}

	updatedUser, err := h.userService.FindByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedUser)
}

func (h *UserHTTPHandler) DeleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return err
	}

	err = h.userService.Delete(id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
