package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	AppPort string
	DbUrl string
}

func (e *EnvConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Env file not loaded!")
	}
	e.AppPort = loadString("APP_PORT", ":3000")
	e.DbUrl = loadString("POSTGRESQL_URL","")

}

var Config EnvConfig

func init() {
	Config.LoadConfig()
}

func loadString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panic("APP_PORT is not loaded")
		return fallback
	}
	return val
}
