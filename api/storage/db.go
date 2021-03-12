package storage

import (
	"echo-rest-api/config"
	"echo-rest-api/models"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	DB, err = gorm.Open(postgres.Open(conString), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	DB.AutoMigrate(&models.User{})

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
