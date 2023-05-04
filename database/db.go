package database

import (
	"github.com/iam-benjamen/Pollie/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
