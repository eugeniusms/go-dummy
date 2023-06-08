package service

import (
	"testing"

	"github.com/eugeniusms/go-dummy/entity"
	"github.com/eugeniusms/go-dummy/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserService{Repository: userRepository}

func TestUserService_GetNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", "1").Return(nil)

	user, err := userService.Get("1")
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_GetSuccess(t *testing.T) {
	user := entity.User{
		Id : "1",
		Name : "John Doe",
		Age : 27,
		Gender : "Male",
		Country : "United Kingdom",
		City : "London",
		Address : "Baker Street",
		Phone : "1234567890",
		Email : "johndoe@gmail.com",
	}
	userRepository.Mock.On("FindById", "1").Return(user)

	result, err := userService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user, result)
}