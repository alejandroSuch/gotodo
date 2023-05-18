package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type SaveConfig func(config Config) error

func SaveYamlConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	cfgFile := getConfigFile()
	err = os.WriteFile(cfgFile, data, 0600)
	if err != nil {
		return err
	}

	fmt.Println("Token saved to", cfgFile)
	return nil
}
