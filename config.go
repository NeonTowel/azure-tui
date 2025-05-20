package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Settings     AzureSettings                `yaml:"settings"`
	Environments []AzureEnvironment           `yaml:"environments"`
	Deployments  map[string][]AzureDeployment `yaml:"deployments"`
}

func loadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func saveConfig(config *Config, filename string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
