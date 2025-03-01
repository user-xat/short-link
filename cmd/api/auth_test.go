package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/auth"
	"github.com/user-xat/short-link/internal/user"
	"github.com/user-xat/short-link/pkg/req"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a2@a.ru",
		Password: "$2a$10$3IIB6Ev.LwyzE9X8sOcWjuazpqxaptrqP8Zr.XQmpQEbX7xwz8iv.",
		Name:     "Test User",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a2@a.ru").
		Delete(&user.User{})
}

func TestLoginSuccess(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)

	conf := configs.LoadApiConfig()
	ts := httptest.NewServer(App(conf))
	defer ts.Close()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "123456",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("status code want %d got %d", http.StatusOK, res.StatusCode)
	}
	logRes, err := req.Decode[auth.LoginResponse](res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if logRes.Token == "" {
		t.Fatalf("token is empty")
	}
}

func TestLoginFailed(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)

	conf := configs.LoadApiConfig()
	ts := httptest.NewServer(App(conf))
	defer ts.Close()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "123",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("status code want %d got %d", http.StatusUnauthorized, res.StatusCode)
	}
}
