package api

import (
	"CRUDUSERS/database"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write json data", "error", err)
		return
	}
}

func NewHandler(db database.Application) http.Handler {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Post("/user", handlerPostUser(db))
	r.Get("/user", handlerGetAllUser(db))
	r.Get("/user/{id}", handlerGetUserByID(db))
	return r
}

func handlerPostUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.Body = http.MaxBytesReader(w, r.Body, 50000)
		var user database.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			slog.Error("failed to decode json data", "error", err)
			sendJSON(
				w,
				Response{Error: "invalid json data"},
				http.StatusBadRequest,
			)
			return
		}

		if err := validateUser(user); len(err) > 0 {
			sendJSON(
				w,
				Response{Error: "invalid user data", Data: err},
				http.StatusBadRequest,
			)
			return
		}
		var id database.ID = database.ID(uuid.New())

		userResponse := db.AddUser(id, user)

		if userResponse.ID == "" {
			sendJSON(
				w,
				Response{Error: "failed to add user"},
				http.StatusInternalServerError,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: userResponse},
			http.StatusCreated,
		)

	}

}

func validateUser(user database.User) []map[string]string {
	errs := []map[string]string{}

	if user.FirstName == "" {
		errs = append(errs, map[string]string{"Nome": "É obrigatório informar o nome"})
	}

	if user.LastName == "" {
		errs = append(errs, map[string]string{"Sobrenome": "É obrigatório informar o sobrenome"})
	}

	return errs
}

func handlerGetUserByID(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			slog.Error("failed to parse user id", "error", err)
			sendJSON(
				w,
				Response{Error: "invalid user id format"},
				http.StatusBadRequest,
			)
			return
		}

		user, found := db.GetUser(database.ID(id))
		if !found {
			sendJSON(
				w,
				Response{Error: "user not found"},
				http.StatusNotFound,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: user},
			http.StatusOK,
		)
	}
}

func handlerGetAllUser(db database.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := db.GetAllUsers()

		if len(users) == 0 {
			sendJSON(
				w,
				Response{Error: "no users found"},
				http.StatusNotFound,
			)
			return
		}

		sendJSON(
			w,
			Response{Data: users},
			http.StatusOK,
		)
	}
}
