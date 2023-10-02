package config

import "github.com/joho/godotenv"

func LoadEnvironmentVariables() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env")
	}
}