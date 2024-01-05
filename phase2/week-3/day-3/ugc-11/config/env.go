package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Value struct {
	Database Database
	Server   Server
	Auth     Auth
}

type Database struct {
	DbUrl      string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

type Server struct {
	Base string
	Port string
}

type Auth struct {
	SecretKey string
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
		Server: Server{
			Base: "",
			Port: os.Getenv("PORT"),
		},
		Auth: Auth{
			SecretKey: os.Getenv("AUTH_SECRETKEY"),
		},
	}, nil
}
