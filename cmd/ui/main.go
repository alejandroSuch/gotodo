package main

import (
	"gotodo/pkg/infrastructure/client/di"
	"gotodo/pkg/ui"
)

func main() {
	win := ui.NewUI(
		di.InitializeListTodosClient(),
		di.InitializeCompleteTodoClient(),
	)
	win.ShowAndRun()
}
