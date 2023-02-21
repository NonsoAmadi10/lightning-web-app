package config

import (
	"github.com/NonsoAmadi10/lightning-web-app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB(model ...interface{}) {
	dbString := utils.GetEnv("DATABASE_URL")
	database, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})

	if err != nil {
		panic("Failed to Connect to Database")
	}

	database.AutoMigrate(model...)

	DB = database
}
