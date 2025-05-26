package config

import (
	"path/filepath"
	"encoding/json"
	"io"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl    		string `json:"db_url"`
	CurrentUserName	string `json:"current_user_name"`
}

func Read() (Config, error){
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) SetUser(name string) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

    c.CurrentUserName = name
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(fullPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
