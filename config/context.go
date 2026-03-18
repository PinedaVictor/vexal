package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Context struct {
	GithubUser string `json:"github_user"`
}

type ContextConfig struct {
	Active   string             `json:"active"`
	Contexts map[string]Context `json:"contexts"`
}

// contextFilePath — resolves ~/.vx/context.json using the OS home dir
func contextFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".vx", "context.json"), nil
}

// LoadContextConfig — reads and unmarshals context.json; returns an empty default if the file doesn't exist yet (first run)
func LoadContextConfig() (ContextConfig, error) {
	empty := ContextConfig{Contexts: make(map[string]Context)}
	path, err := contextFilePath()
	if err != nil {
		return empty, err
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return empty, nil
	}
	if err != nil {
		return empty, err
	}
	var cfg ContextConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return empty, err
	}
	if cfg.Contexts == nil {
		cfg.Contexts = make(map[string]Context)
	}
	return cfg, nil
}

// SaveContextConfig — creates ~/.vx/ if needed, then writes the config as indented JSON with restricted permissions (0600)
func SaveContextConfig(cfg ContextConfig) error {
	path, err := contextFilePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
