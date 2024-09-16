package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DatabaseCredentials *DatabaseCredentials `toml:"database_credentials"`
}

type DatabaseCredentials struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	Port     uint   `toml:"port"`
	Timezone string `toml:"timezone"`
}

func (dc *DatabaseCredentials) Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		dc.Host,
		dc.User,
		dc.Password,
		dc.DBName,
		dc.Port,
		dc.Timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
