package config

type Config struct {
	Port             int
	CorsAllowOrigins string
	RedisConfig      RedisConfig
}

type RedisConfig struct {
	Host     string
	User     string
	Password string
}
