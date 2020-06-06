package config

import (
	"os"
)

const (
	DbHost  = "DB_HOST"
	DbPort  = "DB_PORT"
	DbName  = "DB_NAME"
	DbUser  = "DB_USER"
	DbPass  = "DB_PASS"
	AppPort = "APP_PORT"
	AppHost = "APP_HOST"
	TransportType = "TRANSPORT_TYPE"
	TransportHost = "TRANSPORT_HOST"
	TransportPort = "TRANSPORT_PORT"
	TransportUser = "TRANSPORT_USER"
	TransportPass = "TRANSPORT_PASS"
)

var RequiredConfigFields = []string{
	DbHost,
	DbPort,
	DbName,
	DbUser,
	DbPass,
	AppPort,
	AppHost,
	TransportType,
	TransportHost,
	TransportPort,
	TransportUser,
	TransportPass,
}

type Config struct {
	_map map[string]string
}

func EnvParseToConfigMap() (Config, error) {
	config := Config{_map: make(map[string]string)}

	for _, field := range RequiredConfigFields {
		val := os.Getenv(field)

		if val == "" {
			return config, ErrorConfig("env " + field + " is required")
		}

		config._map[field] = val
	}

	return config, nil
}

func (config *Config) Get(field string) string {
	return config._map[field]
}
