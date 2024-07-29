package server

import (
	"coJournal/internal/server/handler"
	"fmt"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(apiRouter *echo.Group, userHandler *handler.UserHTTPHandler) {
	fmt.Println("Registering user Routes!")
	userRouter := apiRouter.Group("/user")
	userRouter.POST("", userHandler.CreateUser)
	userRouter.GET("", userHandler.GetAllUser)
	userRouter.GET("/:id", userHandler.GetUser)
	userRouter.PUT("/:id", userHandler.UpdateUser)
	userRouter.DELETE("/:id", userHandler.DeleteUser)
}

func RegisterJournalRoutes(apiRouter *echo.Group, userHandler *handler.JournalHTTPHandler) {
}

func RegisterNoteRoutes(apiRouter *echo.Group, NoteHandler *handler.NoteHTTPHandler) {
}
