package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitSql(cfg *Value) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", cfg.Database.DbUrl, cfg.Database.DbPort, cfg.Database.DbUser, cfg.Database.DbPassword, cfg.Database.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			Colorful: true,
			LogLevel: logger.Error,
		}),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
