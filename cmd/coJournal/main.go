package main

import (
	"coJournal/internal/repository/in_memory"
	"coJournal/internal/server/handler"
	"coJournal/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// r := server.SetupRouter()
	// log.Fatal(http.ListenAndServe(":8000", r))

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to CoJournal!")
	})

	// apiGroup := e.Group("/api")
	// apiGroup.GET("", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "what the fuck?")
	// })
	//
	userRepo := in_memory.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHTTPHandler(userService)
	// userGroup := apiGroup.Group("/user")
	// userGroup.POST("", userHandler.CreateUser)
	// userGroup.GET("", userHandler.GetAllUser)
	// userGroup.GET(":id", userHandler.GetUser)
	// userGroup.PUT(":id", userHandler.UpdateUser)
	// userGroup.DELETE(":id", userHandler.DeleteUser)

	e.POST("/api/user", userHandler.CreateUser)
	e.GET("/api/user", userHandler.GetAllUser)
	e.GET("/api/user/:id", userHandler.GetUser)
	e.PUT("/api/user/:id", userHandler.UpdateUser)
	e.DELETE("/api/user/:id", userHandler.DeleteUser)

	e.Logger.Fatal(e.Start(":8000"))
}
