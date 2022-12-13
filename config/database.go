package config

import (
	"os"

	"github.com/joho/godotenv"
)

var DB_USERNAME = ""
var DB_PASSWORD = ""
var DB_HOST = ""
var DB_PORT = ""
var DB_NAME = ""

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
}
