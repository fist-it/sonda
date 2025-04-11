package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"testing"
	"time"
)

// Helper function to write a YAML config file for testing
func writeTestConfig(t *testing.T, fileName string, config *Config) {
	data, err := yaml.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		t.Fatalf("Failed to write test config to file: %v", err)
	}
}

// Test the LoadConfig function
func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name         string
		filePath     string
		config       *Config
		expectedErr  bool
		expectedPort int
	}{
		{
			name:     "Valid config",
			filePath: "valid_config.yaml",
			config: &Config{
				Services: []ServiceConfig{
					{Name: "service1", URL: "http://example.com", Interval: 10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 8080,
			},
			expectedErr:  false,
			expectedPort: 8080,
		},
		{
			name:     "Invalid config (missing service name)",
			filePath: "invalid_config_name.yaml",
			config: &Config{
				Services: []ServiceConfig{
					{URL: "http://example.com", Interval: 10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 8080,
			},
			expectedErr:  true,
			expectedPort: 8080,
		},
		{
			name:     "Invalid config (negative interval)",
			filePath: "invalid_config_interval.yaml",
			config: &Config{
				Services: []ServiceConfig{
					{Name: "service1", URL: "http://example.com", Interval: -10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 8080,
			},
			expectedErr:  true,
			expectedPort: 8080,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write the test configuration file
			writeTestConfig(t, tt.filePath, tt.config)

			// Run the LoadConfig function
			config, err := LoadConfig(tt.filePath)

			// Clean up the file after test
			os.Remove(tt.filePath)

			if tt.expectedErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
				if config.Port != tt.expectedPort {
					t.Errorf("Expected port %d but got %d", tt.expectedPort, config.Port)
				}
			}
		})
	}
}

// Test the config validation function
func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name        string
		config      *Config
		expectedErr bool
	}{
		{
			name: "Valid config",
			config: &Config{
				Services: []ServiceConfig{
					{Name: "service1", URL: "http://example.com", Interval: 10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 8080,
			},
			expectedErr: false,
		},
		{
			name: "Invalid config (empty service name)",
			config: &Config{
				Services: []ServiceConfig{
					{Name: "", URL: "http://example.com", Interval: 10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 8080,
			},
			expectedErr: true,
		},
		{
			name: "Invalid config (empty service URL)",
			config: &Config{
				Services: []ServiceConfig{
					{Name: "service1", URL: "", Interval: 10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 8080,
			},
			expectedErr: true,
		},
		{
			name: "Semi valid config (wrong port)",
			config: &Config{
				Services: []ServiceConfig{
					{Name: "service1", URL: "http://example.com", Interval: 10 * time.Second, Timeout: 5 * time.Second},
				},
				Port: 0,
			},
			expectedErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if tt.expectedErr && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectedErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
