package config

import (
	"fmt"
	"os"
)

var (
	DB_USERNAME = ""
	DB_PASSWORD = ""
	DB_HOST     = ""
	DB_PORT     = ""
	DB_NAME     = ""
)

func InitEnvDB() {
	DB_USERNAME = os.Getenv("MYSQL_USER")
	if DB_USERNAME == "" {
		fmt.Println("DB_USERNAME is empty")
	}

	DB_PASSWORD = os.Getenv("MYSQL_PASSWORD")

	DB_HOST = os.Getenv("MYSQL_HOST")
	if DB_HOST == "" {
		fmt.Println("DB_HOST is empty")
	}

	DB_PORT = os.Getenv("MYSQL_PORT")
	if DB_PORT == "" {
		fmt.Println("DB_PORT is empty")
	}

	DB_NAME = os.Getenv("MYSQL_DBNAME")
	if DB_NAME == "" {
		fmt.Println("DB_NAME is empty")
	}
}

// var (
// 	DB_USERNAME = os.Getenv("MYSQL_USER")
// 	DB_PASSWORD = os.Getenv("MYSQL_PASSWORD")
// 	DB_HOST     = os.Getenv("MYSQL_HOST")
// 	DB_PORT     = os.Getenv("MYSQL_PORT")
// 	DB_NAME     = os.Getenv("MYSQL_DBNAME")
// )

// func init() {
// 	if err := godotenv.Load(".env"); err != nil {
// 		panic(err)
// 	}

// 	DB_USERNAME = os.Getenv("DB_USERNAME")
// 	DB_PASSWORD = os.Getenv("DB_PASSWORD")
// 	DB_HOST = os.Getenv("DB_HOST")
// 	DB_PORT = os.Getenv("DB_PORT")
// 	DB_NAME = os.Getenv("DB_NAME")
// }
