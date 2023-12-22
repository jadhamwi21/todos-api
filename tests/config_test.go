package tests

import (
	"testing"
	"todos-api/config"
	"todos-api/internal/database"

	"github.com/spf13/viper"
)

func TestViperConfig(t *testing.T) {
	config.SetupConfig()

	if len(viper.AllKeys()) == 0 {
		t.Errorf("configuration not set")
	}
}

func TestDatabaseSetup(t *testing.T) {
	config.SetupConfig()
	db, err := database.SetupDatabase()

	defer func() {
		if db != nil {
			db, _ := db.DB()
			db.Close()
		}
	}()
	if err != nil {
		t.Errorf(err.Error())
	}
}
