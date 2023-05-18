package todo

import (
	"errors"
	"github.com/spf13/cobra"
	"gotodo/pkg/infrastructure/config"
)

func NewTodoCommand(cmds ...*cobra.Command) *cobra.Command {
	todoCmd := &cobra.Command{
		Use:          "todo",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.LoadYamlConfig()

			if cfg.Token == "" {
				return errors.New("no token found. Please run 'app login' first")
			}

			return nil
		},
	}

	todoCmd.AddCommand(cmds...)

	return todoCmd
}
