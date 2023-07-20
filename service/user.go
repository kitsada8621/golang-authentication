package service

import (
	"github.com/golang-docker/model"
)

type UserService interface {
	FindOne(field string, params string) (*model.User, error)
	Create(payload *model.User) (*model.User, error)
}
