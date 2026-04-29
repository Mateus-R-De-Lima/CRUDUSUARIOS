package api

import (
	"CRUDUSERS/internal/database"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func handlerDeleteUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			sendJSON(w, Response{Error: "invalid user id format"}, http.StatusBadRequest)
			return
		}

		deleted := db.DeleteUser(database.ID(id))
		if !deleted {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: map[string]string{"message": "user deleted successfully"}}, http.StatusOK)
	}
}
