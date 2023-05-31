package domain

type TodoStatus int

const (
	TodoStatusPending TodoStatus = iota
	TodoStatusCompleted
)

func (s TodoStatus) String() string {
	return [...]string{"Pending", "Completed"}[s]
}

type Todo struct {
	ID          string
	Description string
	Status      TodoStatus
}

func (t *Todo) Complete() {
	t.Status = TodoStatusCompleted
}

type Todos []Todo

func (t Todos) Has(id string) bool {
	for _, it := range t {
		if it.ID == id {
			return true
		}
	}

	return false
}

func (t Todos) Complete(id string) {
	for _, it := range t {
		if it.ID == id {
			it.Complete()
			return
		}
	}
}
