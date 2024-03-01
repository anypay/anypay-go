// Package myapp contains the application logic.
package server

import (
	"log"
	"net/http"

	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

// StartServer starts the Gin server and sets up the routes.
func StartServer(configPath string) {
	router := gin.Default()

	// Load configuration
	config := loadConfig(configPath)

	// add websockets routes if websockets are enabled
	if config.Websockets {
		SetupWebsocketRoutes(router)
	}

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Kubernetes readiness and liveness probes
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	router.GET("/readyz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	// Start the server
	router.Run(":" + config.Port)
}

// loadConfig loads the application configuration from a YAML file.
func loadConfig(path string) *Config {
	defaultConfig := &Config{
		Port:       "8080", // Default port if not specified
		Websockets: bool(true),
	}

	// If path is nil, return the default configuration.
	if path == "" {
		log.Println("No config path provided, using default settings.")
		return defaultConfig
	}

	// Attempt to read the specified config file.
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read config file at %s, using default settings: %v\n", path, err)
		return defaultConfig
	}

	// Attempt to unmarshal the contents of the config file into a Config struct.
	if err := yaml.Unmarshal(yamlFile, defaultConfig); err != nil {
		log.Printf("Failed to unmarshal config, using default settings: %v\n", err)
		return defaultConfig
	}

	return defaultConfig
}
