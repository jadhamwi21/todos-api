package database

import (
	"todos-api/internal/todos"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := viper.Get("DSN").(string)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&todos.Todo{}); err != nil {
		return nil, err
	}

	return db, nil
}
