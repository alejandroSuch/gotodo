package cmd

import "github.com/spf13/cobra"

func NewRootCommand(cmds ...*cobra.Command) *cobra.Command {
	rootCmd := &cobra.Command{Use: "app"}

	rootCmd.AddCommand(cmds...)

	return rootCmd
}
