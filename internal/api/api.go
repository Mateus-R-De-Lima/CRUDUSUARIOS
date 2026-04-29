package api

import (
	"CRUDUSERS/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(db database.Application) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Post("/user", handlerPostUser(db))
	r.Get("/user", handlerGetAllUser(db))
	r.Get("/user/{id}", handlerGetUserByID(db))
	r.Put("/user/{id}", handlerUpdateUser(db))
	r.Delete("/user/{id}", handlerDeleteUser(db))

	return r
}
