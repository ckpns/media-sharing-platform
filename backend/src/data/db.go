package data

import (
	"log"

	"github.com/ckpns/media-sharing-platform/backend/src/data/models"

	"github.com/ckpns/media-sharing-platform/backend/src/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDB() *gorm.DB {

	config := common.GetConfig()

	db, err := gorm.Open(postgres.Open(config.DB.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database", err)
	}

	if config.DB.AutoMigrate {
		db.AutoMigrate(models.User{}, models.Post{}, models.FavoritedPost{})
	}

	database = db

	return database
}

func GetDB() *gorm.DB {
	return database
}
