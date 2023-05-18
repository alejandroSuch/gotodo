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
