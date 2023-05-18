package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LoadConfig func() Config

func LoadYamlConfig() Config {
	var config Config

	data, err := os.ReadFile(getConfigFile())
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return config
	}

	return config
}
