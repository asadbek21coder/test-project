package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	Service_2Host string
	Service_2Port int
	Service_1Host string
	Service_1Port int
	LogLevel      string
	HttpPort      string
}

// Load loads environment vars and inflates Config
func Load() Config {
	

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8090"))

	config.Service_2Host = cast.ToString(getOrReturnDefault("SERVICE_HOST", "localhost"))
	config.Service_2Port = cast.ToInt(getOrReturnDefault("SERVICE_PORT", 9102))
	config.Service_1Host = cast.ToString(getOrReturnDefault("SERVICE_HOST", "localhost"))
	config.Service_1Port = cast.ToInt(getOrReturnDefault("SERVICE_PORT", 9101))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
