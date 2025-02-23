package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/user-xat/short-link/internal/link"
	"github.com/user-xat/short-link/internal/stat"
	"github.com/user-xat/short-link/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
}
