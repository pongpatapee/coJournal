package main

import (
	"coJournal/internal/repository/postgres"
	"coJournal/internal/server"
	"coJournal/internal/server/handler"
	"coJournal/internal/service"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main() {
	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse configuration: %v\n", err)
		os.Exit(1)
	}

	// To disable statement caching
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to CoJournal!")
	})

	apiRouter := e.Group("/api")

	userRepo := postgres.NewPostgresUserRepository(dbpool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHTTPHandler(userService)
	server.RegisterUserRoutes(apiRouter, userHandler)

	journalRepo := postgres.NewPostgresJournalRepository(dbpool)
	journalService := service.NewJournalService(journalRepo)
	journalHandler := handler.NewJournalHTTPHandler(journalService)
	server.RegisterJournalRoutes(apiRouter, journalHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
