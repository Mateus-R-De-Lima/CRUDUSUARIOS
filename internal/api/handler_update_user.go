package api

import (
	"CRUDUSERS/internal/database/store"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlerUpdateUser(queries *store.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := store.ParseUUID(idStr)
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			sendJSON(w, Response{Error: "invalid user id format"}, http.StatusBadRequest)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 50000)
		var user store.UpdateUserParams

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Error("failed to decode json data", "error", err)
			sendJSON(w, Response{Error: "invalid json data"}, http.StatusBadRequest)
			return
		}

		user.ID = id
		if user.ID != id {
			slog.Error("user id in url does not match user id in body", "url_id", id, "body_id", user.ID)
			sendJSON(w, Response{Error: "user id in url does not match user id in body"}, http.StatusBadRequest)
			return
		}

		if err := validateUpdateUser(user); len(err) > 0 {
			sendJSON(w, Response{Error: "invalid user data", Data: err}, http.StatusBadRequest)
			return
		}

		err = queries.UpdateUser(context.Background(), user)

		if err != nil {
			slog.Error("failed to update user", "error", err)
			sendJSON(w, Response{Error: "failed to update user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
