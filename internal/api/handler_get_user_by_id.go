package api

import (
	"CRUDUSERS/internal/database/store"
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlerGetUserByID(queries *store.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := store.ParseUUID(idStr)

		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			sendJSON(w, Response{Error: "invalid user id format"}, http.StatusBadRequest)
			return
		}

		user, err := queries.GetUser(context.Background(), id)
		if err != nil {
			slog.Error("failed to fetch user", "error", err)
			sendJSON(w, Response{Error: "error fetching user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
