package service

import (
	"github.com/golang-docker/dto"
	"github.com/golang-docker/model"
)

type AuthService interface {
	Attempt(credential *dto.Credentials) *model.User
	Login()
	Logout()
	Profile()
	HashedPassword(password string) (string, error)
	ComparePassword(password string, hash string) bool
}
