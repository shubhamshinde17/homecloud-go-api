package config

import (
	"os"
	"strconv"
)

func LoadConfig() *Config {
	jwtExpirationMinutes, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_MINUTES"))
	refreshExpirationDays, _ := strconv.Atoi(os.Getenv("REFRESH_EXPIRATION_DAYS"))

	return &Config{
		Port:                  os.Getenv("PORT"),
		DBHost:                os.Getenv("DB_HOST"),
		DBPort:                os.Getenv("DB_PORT"),
		DBName:                os.Getenv("DB_NAME"),
		DBUser:                os.Getenv("DB_USER"),
		DBPassword:            os.Getenv("DB_PASSWORD"),
		JWTSecret:             os.Getenv("JWT_SECRET"),
		JWTExpirationMinutes:  jwtExpirationMinutes,
		RefreshExpirationDays: refreshExpirationDays,
		StoragePath:           os.Getenv("STORAGE_PATH"),
	}
}
