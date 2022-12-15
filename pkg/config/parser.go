package config

import (
	"os"
	"strconv"
)

const (
	VariablePort             = "PORT"
	VariableCORSAllowOrigins = "CORS_ALLOW_ORIGINS"
)

func ParseFromEnvironmentalVariables() (Config, error) {
	portString := os.Getenv(VariablePort)
	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, err
	}
	corsAllowOrigins := os.Getenv(VariableCORSAllowOrigins)
	return Config{
		Port:             port,
		CorsAllowOrigins: corsAllowOrigins,
	}, nil
}
