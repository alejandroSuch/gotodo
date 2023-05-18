package client

import (
	"fmt"
	"gotodo/pkg/infrastructure/config"
	"net/http"
)

type ListTodos struct {
	url        string
	loadConfig config.LoadConfig
}

func NewListTodosClient(baseUrl string, loadConfig config.LoadConfig) ListTodos {
	return ListTodos{
		url:        fmt.Sprintf("%s/todos", baseUrl),
		loadConfig: loadConfig,
	}
}

func (l ListTodos) Execute() (Todos, error) {
	var todos Todos
	cfg := l.loadConfig()

	err := DoRequest(Request{
		Method:      http.MethodGet,
		Url:         l.url,
		RequestBody: nil,
		ApiResponse: &todos,
		AuthToken:   cfg.Token,
	})

	if err != nil {
		return nil, err
	}

	return todos, nil
}
