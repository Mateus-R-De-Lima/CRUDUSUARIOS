package database

import "github.com/google/uuid"

type ID uuid.UUID

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Biography string `json:"biography"`
}

type UserResponse struct {
	ID   string `json:"id"`
	User User   `json:"user"`
}

type Application struct {
	data map[ID]User
}

func NewApplication() Application {
	return Application{data: make(map[ID]User)}
}

func (app *Application) AddUser(id ID, user User) UserResponse {
	if app.data == nil {
		app.data = make(map[ID]User)
	}
	app.data[id] = user
	return UserResponse{
		ID:   uuid.UUID(id).String(),
		User: user,
	}
}

func (app *Application) GetUser(id ID) (UserResponse, bool) {
	user, ok := app.data[id]
	if !ok {
		return UserResponse{}, false
	}
	return UserResponse{
		ID:   uuid.UUID(id).String(),
		User: user,
	}, true
}

func (app *Application) GetAllUsers() []UserResponse {
	users := make([]UserResponse, 0, len(app.data))
	for id, user := range app.data {
		users = append(users, UserResponse{
			ID:   uuid.UUID(id).String(),
			User: user,
		})
	}
	return users
}
