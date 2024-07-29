package main

import (
	"coJournal/internal/repository/postgres"
	"coJournal/internal/server/handler"
	"coJournal/internal/service"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// r := server.SetupRouter()
	// log.Fatal(http.ListenAndServe(":8000", r))

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

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
	// userRepo := in_memory.NewInMemoryUserRepository()
	userRepo := postgres.NewPostgresUserRepository(dbpool)
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
