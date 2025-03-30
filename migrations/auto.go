package main

import (
	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/models"
	"github.com/user-xat/short-link/internal/stat"
	"github.com/user-xat/short-link/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := configs.LoadApiConfig()
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Link{}, &user.User{}, &stat.Stat{})
}
