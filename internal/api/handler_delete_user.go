package api

import (
	"CRUDUSERS/internal/database/store"
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlerDeleteUser(queries *store.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := store.ParseUUID(idStr)
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			sendJSON(w, Response{Error: "invalid user id format"}, http.StatusBadRequest)
			return
		}

		err = queries.DeleteUser(context.Background(), id)
		if err != nil {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: map[string]string{"message": "user deleted successfully"}}, http.StatusOK)
	}
}
