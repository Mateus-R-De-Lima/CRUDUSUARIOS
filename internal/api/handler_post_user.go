package api

import (
	"CRUDUSERS/internal/database"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

func handlerPostUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		id := database.ID(uuid.New())
		userResponse := db.AddUser(id, user)
		if userResponse.ID == "" {
			sendJSON(w, Response{Error: "failed to add user"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, Response{Data: userResponse}, http.StatusCreated)
	}
}
