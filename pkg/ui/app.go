package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gotodo/pkg/infrastructure/client"
)

func NewUI(listTodos client.ListTodos, completeTodo client.CompleteTodo) fyne.Window {
	app := app.New()
	win := app.NewWindow("Lista de Todos")

	win.SetContent(createUI(listTodos, completeTodo))
	return win
}

func createUI(listTodos client.ListTodos, completeTodo client.CompleteTodo) *fyne.Container {
	vbox := container.NewVBox()

	todos, err := listTodos.Execute()
	if err != nil {
		panic(err)
	}

	for _, todo := range todos {
		if todo.Completed {
			continue
		}

		t := todo
		todoLabel := widget.NewLabel(fmt.Sprintf("%s", todo.Description))

		var hbox *fyne.Container
		completeButton := widget.NewButton("Completar", func() {
			err2 := completeTodo.Execute(client.CompleteTodoCommand{
				TodoId: t.ID,
			})

			if err2 != nil {
				fmt.Println(err)
			}
			vbox.Remove(hbox)
		})

		hbox = container.NewHBox(todoLabel, completeButton)

		vbox.Add(hbox)
	}
	return vbox
}
