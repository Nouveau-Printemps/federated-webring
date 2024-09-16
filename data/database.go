package data

import (
	"github.com/Nouveau-Printemps/federated-webring/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(credentials *config.DatabaseCredentials) {
	db = credentials.Connect()
	err := db.AutoMigrate(&WebsiteType{}, &Website{})
	if err != nil {
		panic(err)
	}

	err = db.Find(&Websites).Preload("type").Error
	if err != nil {
		panic(err)
	}
}
