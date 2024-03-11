package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// GodotenvLoader é uma implementação de EnvLoader que usa godotenv.
type GodotenvLoader struct{}

// Load carrega as variáveis de ambiente usando godotenv.
func (g *GodotenvLoader) Load() error {
	return godotenv.Load()
}

// Configuration armazena as configurações da aplicação.
type Configuration struct {
	RedisAddress       string
	IPRateLimit        int
	TokenRateLimit     int
	TokenBlockDuration int
	IPBlockDuration    int
}

// LoadConfig carrega as configurações a partir de variáveis de ambiente.
func LoadConfig(loader EnvLoader) (*Configuration, error) {
	if err := loader.Load(); err != nil {
		return nil, err
	}

	config := &Configuration{
		RedisAddress:       getEnv("REDIS_ADDRESS", "localhost:6379"),
		IPRateLimit:        getEnvAsInt("IP_RATE_LIMIT", 5),
		TokenRateLimit:     getEnvAsInt("TOKEN_RATE_LIMIT", 100),
		TokenBlockDuration: getEnvAsInt("TOKEN_BLOCK_DURATION", 300),
		IPBlockDuration:    getEnvAsInt("IP_BLOCK_DURATION", 300),
	}

	return config, nil
}

// getEnv lê uma variável de ambiente como string, com um valor padrão.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt lê uma variável de ambiente como int, com um valor padrão.
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
