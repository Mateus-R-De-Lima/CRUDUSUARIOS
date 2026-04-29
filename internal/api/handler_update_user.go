package api

import (
	"CRUDUSERS/internal/database"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func handlerUpdateUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			sendJSON(w, Response{Error: "invalid user id format"}, http.StatusBadRequest)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 50000)
		var user database.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Error("failed to decode json data", "error", err)
			sendJSON(w, Response{Error: "invalid json data"}, http.StatusBadRequest)
			return
		}

		if err := validateUser(user); len(err) > 0 {
			sendJSON(w, Response{Error: "invalid user data", Data: err}, http.StatusBadRequest)
			return
		}

		updatedUser, found := db.UpdateUser(database.ID(id), user)
		if !found {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: updatedUser}, http.StatusOK)
	}
}
