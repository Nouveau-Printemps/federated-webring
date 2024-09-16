package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Name                string               `json:"name"`
	Description         string               `json:"description"`
	Host                string               `json:"host"`
	DatabaseCredentials *DatabaseCredentials `toml:"database_credentials"`
	RedisCredentials    *RedisCredentials    `toml:"redis_credentials"`
}

type DatabaseCredentials struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DBName   string `toml:"dbname"`
	Port     uint   `toml:"port"`
	Timezone string `toml:"timezone"`
}

type RedisCredentials struct {
	Host     string `toml:"host"`
	Password string `toml:"password"`
	Port     uint   `toml:"port"`
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

func (rc *RedisCredentials) Connect() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", rc.Host, rc.Port),
		Password: rc.Password, // no password set
		DB:       0,           // use default DB
	})
	return rdb, rdb.Ping(context.Background()).Err()
}
