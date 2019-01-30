package config

import "os"

func setDatabase() {
	config := make(map[string]string)
	config["dbHost"] = os.Getenv("DB_HOST")
	config["dbPort"] = os.Getenv("DB_PORT")
	config["dbUser"] = os.Getenv("DB_USER")
	config["dbPassword"] = os.Getenv("DB_PASSWORD")
	config["dbName"] = os.Getenv("DB_NAME")
	configs["db"] = config
}
