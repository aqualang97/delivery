package config

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/aqualang97/logger/v4"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port                   string
	AccessSecret           string
	RefreshSecret          string
	AccessLifetimeMinutes  int
	RefreshLifetimeMinutes int
	Logger                 *logger.Logger
	Driver                 string
	DataSourceName         string
}

func NewConfig(l *logger.Logger) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	accessLifetimeMinutes, _ := strconv.Atoi(os.Getenv("ACCESS_LIFE_TIME"))
	refreshLifetimeMinutes, _ := strconv.Atoi(os.Getenv("REFRESH_LIFE_TIME"))

	return &Config{
		Port:                   os.Getenv("PORT"),
		AccessSecret:           os.Getenv("ACCESS_SECRET"),
		RefreshSecret:          os.Getenv("REFRESH_SECRET"),
		AccessLifetimeMinutes:  accessLifetimeMinutes,
		RefreshLifetimeMinutes: refreshLifetimeMinutes,
		Logger:                 l,
		Driver:                 os.Getenv("DRIVER"),
		DataSourceName:         os.Getenv("DATA_SOURCE_NAME"),
	}

}
