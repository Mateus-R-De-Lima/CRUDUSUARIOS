package main

import (
	"CRUDUSERS/internal/api"
	"CRUDUSERS/internal/database/store"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {

		slog.Error("failed to run server", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}

func run() error {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Criando uma conexão com o banco de dados PostgreSQL usando pgxpool, utilizando as variáveis de ambiente para configurar a conexão
	db, err := pgxpool.New(context.Background(), fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("GOCRUD_DATABASE_USER"),
		os.Getenv("GOCRUD_DATABASE_PASSWORD"),
		os.Getenv("GOCRUD_DATABASE_HOST"),
		os.Getenv("GOCRUD_DATABASE_PORT"),
		os.Getenv("GOCRUD_DATABASE_NAME")))

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
