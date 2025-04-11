package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ServiceConfig struct {
	Name     string        `yaml:"name"`
	URL      string        `yaml:"url"`
	Interval time.Duration `yaml:"interval"`
	Timeout  time.Duration `yaml:"timeout"`
}

type Config struct {
	Services []ServiceConfig `yaml:"services"`
	Port     int             `yaml:"port"`
}

func (c *Config) validate() error { // {{{
	if c.Port <= 0 || c.Port > 65535 {
		c.Port = 12005
	}
	
	if len(c.Services) == 0 {
		return fmt.Errorf("at least one service must be defined")
	}


	for _, service := range c.Services {
		if service.Name == "" {
			return fmt.Errorf("service name cannot be empty")
		}
		if service.URL == "" {
			return fmt.Errorf("service URL cannot be empty")
		}
		if service.Interval <= 0 {
			return fmt.Errorf("service interval must be positive")
		}
		if service.Timeout <= 0 {
			return fmt.Errorf("service timeout must be positive")
		}
	}

	return nil
}

// }}}

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return &config, nil
}


