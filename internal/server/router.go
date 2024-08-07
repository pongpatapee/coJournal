package server

import (
	"coJournal/internal/server/handler"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(apiRouter *echo.Group, userHandler *handler.UserHTTPHandler) {
	userRouter := apiRouter.Group("/user")
	userRouter.POST("", userHandler.CreateUser)
	userRouter.GET("", userHandler.GetAllUser)
	userRouter.GET("/:id", userHandler.GetUser)
	userRouter.PUT("/:id", userHandler.UpdateUser)
	userRouter.DELETE("/:id", userHandler.DeleteUser)
}

func RegisterJournalRoutes(apiRouter *echo.Group, journalHandler *handler.JournalHTTPHandler) {
	journalRouter := apiRouter.Group("/journal")
	journalRouter.POST("", journalHandler.CreateJournal)
	journalRouter.GET("", journalHandler.GetAllJournal)
	journalRouter.GET("/:id", journalHandler.GetJournal)
	journalRouter.PUT("/:id", journalHandler.UpdateJournal)
	journalRouter.DELETE("/:id", journalHandler.DeleteJournal)
}

func RegisterNoteRoutes(apiRouter *echo.Group, noteHandler *handler.NoteHTTPHandler) {
	noteRouter := apiRouter.Group("/note")
	noteRouter.POST("", noteHandler.CreateNote)
	noteRouter.GET("", noteHandler.GetAllNote)
	noteRouter.GET("/:id", noteHandler.GetNote)
	noteRouter.PUT("/:id", noteHandler.UpdateNote)
	noteRouter.DELETE("/:id", noteHandler.DeleteNote)
}
