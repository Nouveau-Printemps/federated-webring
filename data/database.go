package data

import (
	"github.com/Nouveau-Printemps/federated-webring/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var db *gorm.DB

var rdb *redis.Client

func Init(pgc *config.DatabaseCredentials, rdc *config.RedisCredentials) {
	db = pgc.Connect()
	err := db.AutoMigrate(&WebsiteType{}, &Website{})
	if err != nil {
		panic(err)
	}

	err = db.Find(&Websites).Preload("type").Error
	if err != nil {
		panic(err)
	}

	rdb, err = rdc.Connect()
	if err != nil {
		panic(err)
	}
}
