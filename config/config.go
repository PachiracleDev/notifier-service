package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App    AppConfig
	Server ServerConfig

	AWS  AWSConfig
	Auth AuthConfig
}

type AppConfig struct {
	Name    string
	Version string
	Env     string
}

type ServerConfig struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AWSConfig struct {
	Region          string
	AccessKeyId     string
	SecretAccessKey string
	BucketPublic    string
	BucketPrivate   string
}

type AuthConfig struct {
	Username string
	Password string
}

func LoadConfig(dir string) {
	err := godotenv.Load(dir)
	if err != nil {
		log.Printf("No se pudo cargar el archivo .env, usando variables de entorno del sistema: %v", err)
	}
}

func NewConfig() *Config {
	LoadConfig(".env")
	config := &Config{
		App: AppConfig{
			Name:    os.Getenv("APP_NAME"),
			Version: os.Getenv("APP_VERSION"),
			Env:     os.Getenv("APP_ENV"),
		},
		Server: ServerConfig{
			Port:         getEnvAsInt("SERVER_PORT", 8080),
			ReadTimeout:  getEnvAsDuration("SERVER_READ_TIMEOUT", 10) * time.Second,
			WriteTimeout: getEnvAsDuration("SERVER_WRITE_TIMEOUT", 10) * time.Second,
		},
		AWS: AWSConfig{
			Region:          os.Getenv("AWS_REGION"),
			AccessKeyId:     os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		},
		Auth: AuthConfig{
			Username: os.Getenv("AUTH_USERNAME"),
			Password: os.Getenv("AUTH_PASSWORD"),
		},
	}

	return config
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsDuration(name string, defaultVal int) time.Duration {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return time.Duration(value)
	}
	return time.Duration(defaultVal)
}
