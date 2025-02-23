package jwt_test

import (
	"testing"

	"github.com/user-xat/short-link/pkg/jwt"
)

func TestJWTCreate(t *testing.T) {
	const email = "a@a.ru"
	jwtService := jwt.NewJWT("my-test-secret")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	data, ok := jwtService.Parse(token)
	if !ok {
		t.Fatalf("unable parse token")
	}
	if data.Email != email {
		t.Fatalf("email want %s got %s", email, data.Email)
	}
}
