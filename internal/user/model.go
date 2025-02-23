package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func NewUser(email, pass, name string) *User {
	return &User{
		Email:    email,
		Password: pass,
		Name:     name,
	}
}
