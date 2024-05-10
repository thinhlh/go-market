package database

import (
	"github.com/thinhlh/go-market/internal/app/product/domain"
)

func (db *Database) Migration() error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	return db.AutoMigrate(&domain.Product{})
}
