package helper

import (
	"testing"

	go_dummy "github.com/eugeniusms/go-dummy"
)

func TestGetUserName(t *testing.T) {
	result := go_dummy.GetUser()["name"]
	if result != "John Do" {
		t.Fatalf("Expected John Doe, got %s", result)
	}
}