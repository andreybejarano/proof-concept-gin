package config

import (
	"log"

	"github.com/joho/godotenv"
)

var configs = make(map[string]map[string]string)

// Init : config initial
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	setServer()
	setDatabase()
}

// GetConfig : return all configs
func GetConfig() map[string]map[string]string {
	return configs
}
