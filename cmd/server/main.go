package main

import (
	"gotodo/pkg/infrastructure/server/db"
	"gotodo/pkg/infrastructure/server/di"
	"gotodo/pkg/rest"
)

func main() {
	db.MigrateDb()

	rest.Start(
		di.InitializeLoginController(),
		di.InitializeRegisterController(),
		di.InitializeCreateTodoController(),
		di.InitializeCompleteTodoController(),
		di.InitializeListTodosController(),
		di.InitializeAuthMiddleware(),
	)
}
