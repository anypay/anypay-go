// Package myapp contains the core logic and structures of the application.
package server

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config represents the structure of the configuration file.
// Add other configuration parameters as needed.
type Config struct {
	Port       string `yaml:"port"`
	Websockets bool   `yaml:"websockets"`
}

// LoadConfig reads a YAML-formatted configuration file and unmarshals it into a Config struct.
func LoadConfig(configPath string) *Config {
	// Read the file from the given path
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Unable to read config file: %v", err)
	}

	// Unmarshal the YAML into a Config struct
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}

	return &config
}
