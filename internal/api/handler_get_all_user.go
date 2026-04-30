package api

import (
	"CRUDUSERS/internal/database/store"
	"context"
	"net/http"
)

func handlerGetAllUser(queries *store.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contexto := context.Background()

		users, err := queries.ListUsers(contexto)
		if err != nil {
			sendJSON(w, Response{Error: "error fetching users"}, http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			sendJSON(w, Response{Error: "no users found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: users}, http.StatusOK)
	}
}
