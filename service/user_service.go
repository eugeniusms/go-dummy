package service

import (
	"errors"

	"github.com/eugeniusms/go-dummy/entity"
	"github.com/eugeniusms/go-dummy/repository"
)

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) Get(id string) (*entity.User, error) {
	user := service.Repository.FindById(id)
	if user == nil {
		return user, errors.New("User not found")
	} else {
		return user, nil
	}
}