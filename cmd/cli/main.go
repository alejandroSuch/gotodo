package main

import (
	"fmt"
	"gotodo/pkg/infrastructure/cmd/di"
	"os"
)

func main() {
	rootCmd := di.InitializeRootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
