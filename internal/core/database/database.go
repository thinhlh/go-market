package database

import (
	"fmt"

	"github.com/thinhlh/go-market/internal/core/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct{ *gorm.DB }

func NewDatabaseConnection(config config.DatabaseConfig) (*Database, error) {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	dns := fmt.Sprintf(
		"host=%v user=%v  password=%v dbname=%v port=%v",
		config.DatabaseHost,
		config.DatabaseUsername,
		config.DatabasePassword,
		config.DatabaseName,
		config.DatabasePort,
	)

	connection, err := gorm.Open(postgres.Open(dns), gormConfig)

	return &Database{connection}, err
}
