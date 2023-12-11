package entity

import (
	"github.com/rafaelsouzaribeiro/9-API/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id    entity.Id `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	// Nunca sera exibido
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	return &User{
		Id:       entity.NewId(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
