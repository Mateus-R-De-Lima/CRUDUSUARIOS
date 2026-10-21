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
	r.Route("/user", func(r chi.Router) {
		r.Use(JWTMiddleware)
		r.Get("/", handlerGetAllUser(queries))
		r.Get("/{id}", handlerGetUserByID(queries))
		r.Put("/{id}", handlerUpdateUser(queries))
		r.Delete("/{id}", handlerDeleteUser(queries))

	})

	r.Post("/login", LoginHandler)
	r.Post("/refresh", RefreshHandler)

	return r
}
