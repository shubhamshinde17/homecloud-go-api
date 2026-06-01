package config

type Config struct {
	Port                  string
	DBHost                string
	DBPort                string
	DBName                string
	DBUser                string
	DBPassword            string
	JWTSecret             string
	JWTExpirationMinutes  int
	RefreshExpirationDays int
	StoragePath           string
}
