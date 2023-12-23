package database

import (
	"fmt"
	"todos-api/internal/auth"
	"todos-api/internal/todos"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	fmt.Println(viper.Get("DSN"))
	dsn := viper.Get("DSN").(string)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&todos.Todo{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&auth.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
