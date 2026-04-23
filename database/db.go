package database

import "github.com/google/uuid"

type ID uuid.UUID

type User struct {
	FirstName string
	LastName  string
	Biography string
}

type Application struct {
	data map[ID]User
}

func NewApplication() Application {
	return Application{data: make(map[ID]User)}
}

func (app *Application) AddUser(id ID, user User) {
	if app.data == nil {
		app.data = make(map[ID]User)
	}
	app.data[id] = user
}
