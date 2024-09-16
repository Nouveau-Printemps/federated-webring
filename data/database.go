package data

import (
	"github.com/Nouveau-Printemps/federated-webring/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(credentials *config.DatabaseCredentials) {
	db = credentials.Connect()
	err := db.AutoMigrate()
	if err != nil {
		panic(err)
	}
}
