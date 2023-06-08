package helper

import (
	"testing"

	go_dummy "github.com/eugeniusms/go-dummy"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	test := map[string]interface{}{
		"name": "John Doe",
		"age": 27,
		"gender": "Male",
		"country": "United Kingdom",
		"city": "London",
		"address": "Baker Street",
		"phone": "1234567890",
		"email": "johndoe@gmail.com",
	}
	result := go_dummy.GetUser()
	assert.Equal(t, test, result, "The two maps should be the same.")
}