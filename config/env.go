package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv carrega as variáveis de ambiente
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, usando variáveis de ambiente padrão")
	}
}

// GetEnv pega uma variável de ambiente, se não existir retorna um valor padrão
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
