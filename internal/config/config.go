package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort string
	Env     string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret         string
	AccessTokenTTLMin int
}

func LoadConfig() *Config {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env dosyası bulunamadı, environment değişkenlerinden devam ediliyor.")
	}

	cfg := &Config{
		AppName:    getEnv("APP_NAME", "BackendProject"),
		AppPort:    getEnv("APP_PORT", "8080"),
		Env:        getEnv("ENV", "development"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "backend_db"),

		JWTSecret:         getEnv("JWT_SECRET", "devsecret-change-me"),
		AccessTokenTTLMin: mustAtoi(getEnv("ACCESS_TOKEN_TTL_MIN", "60")),
	}

	return cfg
}
func mustAtoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
