package api

import (
	"CRUDUSERS/internal/database/store"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

func handlerPostUser(queries *store.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 50000)

		var user store.CreateUserParams
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Error("failed to decode json data", "error", err)
			sendJSON(w, Response{Error: "invalid json data"}, http.StatusBadRequest)
			return
		}

		if err := validateCreateUser(user); len(err) > 0 {
			sendJSON(w, Response{Error: "invalid user data", Data: err}, http.StatusBadRequest)
			return
		}

		id, err := store.NewUUID()

		user.ID = id
		if err != nil {
			slog.Error("failed to generate user id", "error", err)
			sendJSON(w, Response{Error: "failed to generate user id"}, http.StatusInternalServerError)
			return
		}

		err = queries.CreateUser(context.Background(), user)

		if err != nil {
			slog.Error("failed to add user", "error", err)
			sendJSON(w, Response{Error: "failed to add user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: user}, http.StatusCreated)
	}
}
