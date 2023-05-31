package domain

import "testing"

func TestTodos_Has(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		t    Todos
		args args
		want bool
	}{
		{
			name: "it should return true when item exists",
			t: Todos{
				Todo{ID: "one", Description: "todo number one"},
			},
			args: args{id: "one"},
			want: true,
		},
		{
			name: "it should return false when item doesn't exist",
			t: Todos{
				Todo{ID: "one", Description: "todo number one"},
			},
			args: args{id: "two"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Has(tt.args.id); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodo_Initial_status(t *testing.T) {
	todo := Todo{}

	if todo.Status != TodoStatusPending {
		t.Errorf("status initial value is not TodoStatusPending")
	}
}

func TestTodo_Complete(t *testing.T) {
	todo := Todo{}
	todo.Complete()

	if todo.Status != TodoStatusCompleted {
		t.Errorf("status hasn't changed to complete")
	}
}
