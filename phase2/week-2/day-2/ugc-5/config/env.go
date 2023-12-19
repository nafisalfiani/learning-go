package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Value struct {
	Database Database
}

type Database struct {
	DbUrl      string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func InitEnv() (*Value, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Value{
		Database: Database{
			DbUrl:      os.Getenv("DB_URL"),
			DbPort:     os.Getenv("DB_PORT"),
			DbName:     os.Getenv("DB_NAME"),
			DbUser:     os.Getenv("DB_USER"),
			DbPassword: os.Getenv("DB_PASSWORD"),
		},
	}, nil
}
