package domain

import "errors"

var (
	ErrTodoNotFound = errors.New("todo not found")
)

type User struct {
	ID     string
	AuthId string
	Name   string
	Todos  []Todo
}

func (u *User) AddTodo(description string) {
	u.Todos = append(u.Todos, Todo{Description: description})
}

func (u *User) CompleteTodo(id string) error {
	for i, todo := range u.Todos {
		if todo.ID == id {
			u.Todos[i].Status = TodoStatusCompleted
		}
	}

	return ErrTodoNotFound
}

func NewUser(id string, authId string, name string) User {
	return User{
		ID:     id,
		AuthId: authId,
		Name:   name,
		Todos:  []Todo{},
	}
}
