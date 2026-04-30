package api

import (
	"CRUDUSERS/internal/database/store"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(queries *store.Queries) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Post("/user", handlerPostUser(queries))
	r.Get("/user", handlerGetAllUser(queries))
	r.Get("/user/{id}", handlerGetUserByID(queries))
	r.Put("/user/{id}", handlerUpdateUser(queries))
	r.Delete("/user/{id}", handlerDeleteUser(queries))

	return r
}
