package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getKeyEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Necessario cria a chave '%s'", key)
	}

	return value
}

func EnvConfigDBUrl() string {
	if err := godotenv.Load(); err != nil {
		log.Printf("Faltou o arquivo .env")
	}

	DBHost := getKeyEnv("POSTGRES_HOST")
	DBUser := getKeyEnv("POSTGRES_USER")
	DBName := getKeyEnv("POSTGRES_DB")
	DBPort := getKeyEnv("POSTGRES_PORT")
	DBPassword := getKeyEnv("POSTGRES_PASSWORD")

	DBUrl := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)

	return DBUrl
}
