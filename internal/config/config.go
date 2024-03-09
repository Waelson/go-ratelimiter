package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config armazena as configurações da aplicação
type Configuration struct {
	RedisAddress   string
	IPRateLimit    int
	TokenRateLimit int
	BlockDuration  int
}

// LoadConfig carrega as configurações a partir de variáveis de ambiente
func LoadConfig() Configuration {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := Configuration{
		RedisAddress:   getEnv("REDIS_ADDRESS", "localhost:6379"),
		IPRateLimit:    getEnvAsInt("IP_RATE_LIMIT", 5),
		TokenRateLimit: getEnvAsInt("TOKEN_RATE_LIMIT", 100),
		BlockDuration:  getEnvAsInt("BLOCK_DURATION", 300),
	}

	return config
}

// getEnv lê uma variável de ambiente como string, com um valor padrão
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt lê uma variável de ambiente como int, com um valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
