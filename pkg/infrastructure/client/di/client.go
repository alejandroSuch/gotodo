package di

import (
	"gotodo/pkg/infrastructure/client"
	"gotodo/pkg/infrastructure/config"
)

func InitializeLoginClient() client.Login {
	return client.NewLoginClient(baseUrl, config.SaveYamlConfig)
}

func InitializeListTodosClient() client.ListTodos {
	return client.NewListTodosClient(baseUrl, config.LoadYamlConfig)
}

func InitializeCompleteTodoClient() client.CompleteTodo {
	return client.NewCompleteTodoClient(baseUrl, config.LoadYamlConfig)
}
