package repository

import "github.com/eugeniusms/go-dummy/entity"

type UserRepository interface {
	FindById(id string) *entity.User
}