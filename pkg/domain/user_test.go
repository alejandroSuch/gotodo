package domain

import (
	"testing"
)

func TestUser_AddTodo(t *testing.T) {
	u := &User{Name: "Bruce Wayne"}

	u.AddTodo("I am a todo")

	if len(u.Todos) == 0 {
		t.Errorf("We should have 1 todo\n")
	}

	if u.Todos[0].Description != "I am a todo" {
		t.Errorf("Expected \"I am a todo\". Got %s.\n", u.Todos[0].Description)
	}

	if u.Todos[0].Status != TodoStatusPending {
		t.Errorf("Unexpected status %d.\n", u.Todos[0].Status)
	}
}

func TestUser_CompleteTodo(t *testing.T) {
	type fields struct {
		ID     string
		AuthId string
		Name   string
		Todos  Todos
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "it should error when trying to complete a todo that does not exist",
			fields: fields{
				Name:  "Clark Kent",
				Todos: Todos{},
			},
			args:    args{id: "some-id"},
			wantErr: true,
		},
		{
			name: "it should not error when trying to complete a todo that exists",
			fields: fields{
				Name:  "Clark Kent",
				Todos: Todos{Todo{ID: "some-id"}},
			},
			args:    args{id: "some-id"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Name:  tt.fields.Name,
				Todos: tt.fields.Todos,
			}

			if err := u.CompleteTodo(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CompleteTodo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
