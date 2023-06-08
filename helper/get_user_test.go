package helper

import (
	"testing"

	go_dummy "github.com/eugeniusms/go-dummy"
	"github.com/stretchr/testify/assert"
)

func TestGetUserName(t *testing.T) {
	result := go_dummy.GetUser()["name"]
	assert.Equal(t, "John Doe", result, "The name should be John Doe")
}