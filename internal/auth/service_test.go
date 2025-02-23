package auth_test

import (
	"fmt"
	"testing"

	"github.com/user-xat/short-link/internal/auth"
	"github.com/user-xat/short-link/internal/user"
)

type MockUserRepository map[string]*user.User

func (r MockUserRepository) Create(u *user.User) (*user.User, error) {
	if _, ok := r[u.Email]; ok {
		return nil, fmt.Errorf("user %s already exists", u.Email)
	}
	r[u.Email] = u
	return u, nil
}

func (r MockUserRepository) FindByEmail(email string) (*user.User, error) {
	found, ok := r[email]
	if !ok {
		return nil, fmt.Errorf("user with email %s does not found", email)
	}
	return found, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.ru"
	authService := auth.NewAuthService(MockUserRepository{})
	email, err := authService.Register(initialEmail, "12345", "Jason")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("email %s do not match %s", email, initialEmail)
	}
}
