package todo

import (
	"fmt"
	"github.com/spf13/cobra"
	"gotodo/pkg/infrastructure/client"
)

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

func NewTodoListCommand(client client.ListTodos) *cobra.Command {
	listCmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			todos, err := client.Execute()

			if err != nil {
				fmt.Println("Error al obtener el listado de tareas:", err)
				return
			}

			fmt.Println("Here's your todo list")
			fmt.Println("ID\t\t\t\t\tStatus\tDescription")
			for _, t := range todos {
				fmt.Printf("%s\t%s\t%s\n", t.ID, t.IsCompleted(), t.Description)
			}
		},
	}

	return listCmd
}
