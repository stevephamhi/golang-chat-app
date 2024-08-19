package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BackoffieStaticDir  string
	BackofficeServePort string
	DBDriver            string
	DBUser              string
	DBPassword          string
	DBAddress           string
	DBName              string
	DBNet               string
}

var Env = InitConfig()

func InitConfig() *Config {
	godotenv.Load()

	return &Config{
		BackoffieStaticDir:  GetEnv("BACKOFFICE_STATIC_DIR"),
		BackofficeServePort: GetEnv("BACKOFFICE_SERVE_PORT"),
		DBDriver:            GetEnv("DB_DRIVER", "mysql"),
		DBUser:              GetEnv("DB_USERNAME"),
		DBPassword:          GetEnv("DB_PASSWORD"),
		DBAddress:           fmt.Sprintf("%s:%s", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306")),
		DBName:              GetEnv("DB_NAME", "godbinit"),
		DBNet:               GetEnv("DB_NET", "tcp"),
	}
}

// Get the env by key or fallbacks
func GetEnv(key string, fallback ...string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	if len(fallback) > 0 {
		return fallback[0]
	}

	return ""
}
