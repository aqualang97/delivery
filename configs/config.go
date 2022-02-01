package config

import (
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
}

func NewConfig() *Config {
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
	}
}
