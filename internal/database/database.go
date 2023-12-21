package database

import (
	"todos-api/internal/todos"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := viper.Get("DSN").(string)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&todos.Todo{})

	return db
}
