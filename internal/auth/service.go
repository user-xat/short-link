package auth

import (
	"errors"

	"github.com/user-xat/short-link/internal/user"
	"github.com/user-xat/short-link/pkg/di"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository di.IUserRepository
}

func NewAuthService(userRepository di.IUserRepository) *AuthService {
	return &AuthService{UserRepository: userRepository}
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(hashedPass),
		Name:     name,
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New(ErrWrongCredetials)
	}
	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrWrongCredetials)
	}
	return existedUser.Email, nil
}
