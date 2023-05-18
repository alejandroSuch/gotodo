package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gotodo/pkg/infrastructure/config"
	"net/http"
)

func NewCreateTodoCommand(baseUrl string) *cobra.Command {
	createUrl := fmt.Sprintf("%s/todo", baseUrl)

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new todo",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.LoadYamlConfig()
			description, _ := cmd.Flags().GetString("description")

			payload, _ := json.Marshal(map[string]interface{}{
				"description": description,
			})

			req, err := http.NewRequest("POST", createUrl, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error al crear la solicitud:", err)
				return
			}

			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+cfg.Token)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error al realizar la solicitud:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				fmt.Println("Error en la respuesta del servidor:", resp.Status)
				return
			}

			fmt.Println("Tarea creada. Recuerda lanzar el comando 'app todo list' para ver tu lista de tareas")
		},
	}

	cmd.Flags().String("description", "", "Description")
	cmd.MarkFlagRequired("description")

	return cmd
}
