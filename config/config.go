package config

import (
	"os"
	"strconv"
)

//Config hghg
type Config struct {
	DB   *DBConfig
	PORT string
}

//DBConfig kf
type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

//GetConfig aff
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  os.Getenv("DIALECT"),
			Host:     os.Getenv("HOST"),
			Port:     getEnvAsInt("DBPORT", 3000),
			User:     os.Getenv("DBUSER"),
			Password: os.Getenv("PASSWORD"),
			Dbname:   os.Getenv("DBNAME"),
		},
		PORT: os.Getenv("PORT"),
	}
}

// postgres://zelmmuwe:kg7hf-hRvWz1Szylce4pFDvTDfBuAZDq@raja.db.elephantsql.com:5432/zelmmuwe
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
