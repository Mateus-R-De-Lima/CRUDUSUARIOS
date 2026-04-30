package main

import (
	"CRUDUSERS/internal/api"
	"CRUDUSERS/internal/database/store"
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	if err := run(); err != nil {

		slog.Error("failed to run server", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}

func run() error {
	// Example connection URL: "postgres://username:password@localhost:5432/database_name"
	urlExample := "postgres://pguser:pgpassword@localhost:5432/pguserdb"

	db, err := pgxpool.New(context.Background(), urlExample)

	if err != nil {
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	defer db.Close()

	queries := store.New(db)
	handler := api.NewHandler(queries)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
