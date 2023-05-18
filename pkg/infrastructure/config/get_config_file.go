package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func getConfigFile() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return filepath.Join(home, ".app.yaml")
}
