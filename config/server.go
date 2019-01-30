package config

import "os"

func setServer() {
	config := make(map[string]string)
	config["port"] = os.Getenv("PORT")
	configs["server"] = config
}
