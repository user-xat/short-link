package auth_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/internal/auth"
	"github.com/user-xat/short-link/internal/user"
	"github.com/user-xat/short-link/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("failed init mock db: %v", err)
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}))
	if err != nil {
		return nil, nil, fmt.Errorf("failed init gorm db: %v", err)
	}
	userRepo := user.NewUserRepository(&db.Db{
		DB: gormDb,
	})

	handler := auth.AuthHandler{
		ApiConfig: &configs.ApiConfig{
			Auth: configs.AuthConfig{
				Secret: "secret",
			},
		},
		AuthService: auth.NewAuthService(userRepo),
	}
	return &handler, mock, nil
}

func TestLoginHandlerSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}
	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("a@a.ru", "$2a$10$2Lr2uCPYrfymTtiNQEgbLeYYnnteA49I7PQAnyW.TcuLBElYDp2hG")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a@a.ru",
		Password: "1",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got %d want %d", http.StatusOK, wr.Code)
	}
}

func TestRegisterHandlerSuccess2(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatal(err)
	}
	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	data, _ := json.Marshal(&auth.RegisterRequest{
		Email:    "a@a.ru",
		Password: "1",
		Name:     "Jason",
	})
	reader := bytes.NewReader(data)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(w, r)
	if w.Code != http.StatusCreated {
		t.Fatalf("got %d want %d", w.Code, http.StatusCreated)
	}
}
