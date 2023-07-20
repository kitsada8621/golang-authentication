package service

import (
	"github.com/golang-docker/dto"
	"github.com/golang-docker/model"
	"golang.org/x/crypto/bcrypt"
)

var (
	userManager UserService = NewUserService()
)

type AuthServiceImplement struct {
}

func NewAuthService() AuthService {
	return &AuthServiceImplement{}
}

func (auth *AuthServiceImplement) Attempt(credential *dto.Credentials) *model.User {
	user, err := userManager.FindOne("username", credential.Username)
	if err != nil {
		return nil
	}

	if !auth.ComparePassword(credential.Password, user.Password) {
		return nil
	}

	return user
}

func (auth *AuthServiceImplement) Login() {

}

func (auth *AuthServiceImplement) Profile() {

}

func (auth *AuthServiceImplement) Logout() {

}

func (auth *AuthServiceImplement) HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (auth *AuthServiceImplement) ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
