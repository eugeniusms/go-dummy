package helper

import (
	"testing"

	go_dummy "github.com/eugeniusms/go-dummy"
)

func TestGetUserName(t *testing.T) {
	result := go_dummy.GetUser()["name"]
	if result != "John Doe" {
		t.Errorf("GetUser() failed, expected %v, got %v", "John Doe", result)
		panic("GetUser() failed, expected John Doe")
	}
}