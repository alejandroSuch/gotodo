package client

type Todo struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (t Todo) IsCompleted() string {
	if t.Completed {
		return "👍"
	}

	return "👎"
}

type Todos []Todo
