package client

import (
	"fmt"
	"gotodo/pkg/infrastructure/config"
	"net/http"
	"strings"
)

type CompleteTodo struct {
	url        string
	loadConfig config.LoadConfig
}

type CompleteTodoCommand struct {
	TodoId string
}

func NewCompleteTodoClient(
	baseUrl string,
	loadConfig config.LoadConfig,
) CompleteTodo {
	return CompleteTodo{
		url:        fmt.Sprintf("%s/todos/:id/complete", baseUrl),
		loadConfig: loadConfig,
	}
}

func (ct CompleteTodo) Execute(c CompleteTodoCommand) error {
	cfg := ct.loadConfig()

	err := DoRequest(Request{
		Method:      http.MethodPatch,
		Url:         strings.Replace(ct.url, ":id", c.TodoId, -1),
		RequestBody: nil,
		ApiResponse: nil,
		AuthToken:   cfg.Token,
	})

	if err != nil {
		return err
	}

	return nil
}
