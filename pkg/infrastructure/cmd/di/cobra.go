package di

import (
	"github.com/spf13/cobra"
	"gotodo/pkg/cmd"
	"gotodo/pkg/cmd/auth"
	"gotodo/pkg/cmd/todo"
	"gotodo/pkg/infrastructure/client/di"
)

func InitializeCreateTodoCommand() *cobra.Command {
	c := todo.NewCreateTodoCommand(baseUrl)
	return c
}

func InitializeTodoListCommand() *cobra.Command {
	c := todo.NewTodoListCommand(di.InitializeListTodosClient())
	return c
}

func InitializeTodoCommand() *cobra.Command {
	c := todo.NewTodoCommand(
		InitializeTodoListCommand(),
		InitializeCreateTodoCommand(),
	)
	return c
}

func InitializeLoginCommand() *cobra.Command {
	c := auth.NewLoginCommand(di.InitializeLoginClient())
	return c
}

func InitializeRootCommand() *cobra.Command {
	c := cmd.NewRootCommand(
		InitializeLoginCommand(),
		InitializeTodoCommand(),
	)
	return c
}
