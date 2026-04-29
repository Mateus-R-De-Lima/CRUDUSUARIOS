package api

import (
	"CRUDUSERS/internal/database"
	"net/http"
)

func handlerGetAllUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := db.GetAllUsers()
		if len(users) == 0 {
			sendJSON(w, Response{Error: "no users found"}, http.StatusNotFound)
			return
		}

		sendJSON(w, Response{Data: users}, http.StatusOK)
	}
}
