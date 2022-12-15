package config

import (
	"os"
	"strconv"
)

const (
	VariablePort             = "PORT"
	VariableCORSAllowOrigins = "CORS_ALLOW_ORIGINS"
	VariableRedisHost        = "REDIS_HOST"
	VariableRedisUser        = "REDIS_USER"
	VariableRedisPassword    = "REDIS_PASSWORD"
)

func ParseFromEnvironmentalVariables() (Config, error) {
	portString := os.Getenv(VariablePort)
	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, err
	}
	corsAllowOrigins := os.Getenv(VariableCORSAllowOrigins)
	redisHost := os.Getenv(VariableRedisHost)
	redisUser := os.Getenv(VariableRedisUser)
	redisPassword := os.Getenv(VariableRedisPassword)
	return Config{
		Port:             port,
		CorsAllowOrigins: corsAllowOrigins,
		RedisConfig: RedisConfig{
			Host:     redisHost,
			User:     redisUser,
			Password: redisPassword,
		},
	}, nil
}
