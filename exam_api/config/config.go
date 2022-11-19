package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production
	//custumer
	CustumerServiceHost string
	CustumerServicePort int
	//post
	PostServiceHost string
	PostServicePort int
	//reating5432
	ReatingServiceHost string
	ReatingServicePort int
	// context timeout in seconds
	CtxTimeout int

	RedisHost string
	RedisPort string

	PostgresHost string
	PostgresPort int
	PostgresUser string
	PostgresPassword string
	PostgresDB string


	SignKey string
	AuthConfigPath string
	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":9079"))

	c.CustumerServiceHost = cast.ToString(getOrReturnDefault("CUSTUMER_SERVICE_HOST", "localhost"))
	c.CustumerServicePort = cast.ToInt(getOrReturnDefault("CUSTUMER_SERVICE_PORT", 9088))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "post_service"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVIC_PORT", 9093))

	c.ReatingServiceHost = cast.ToString(getOrReturnDefault("REATING_SERVICE_HOST", "reating_service"))
	c.ReatingServicePort = cast.ToInt(getOrReturnDefault("RATING_SERVIC_PORT", 9084))

	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST","redis"))
	c.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT","6379"))

	c.SignKey = cast.ToString(getOrReturnDefault("SIGN_KEY", "secret"))
	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_PATH", "./config/rabc_model.conf"))
	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "database-1.c9lxq3r1itbt.us-east-1.rds.amazonaws.com"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "rolesdb"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "abduazim"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))


	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
