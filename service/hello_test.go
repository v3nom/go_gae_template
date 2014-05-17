package service

import (
	"testing"
)

func TestgetUsername(t *testing.T) {
	// same package so we can call private methods
	username := getUsername()
	expectedUsername := "User"
	if username != expectedUsername {
		// By default test passes, error is a special case
		t.Errorf("getUsername() = %+v, expected %+v", username, expectedUsername)
	}
}

func TestGetGreeting(t *testing.T) {
	r := GetGreeting()
	expectedGreeting := "Hello, " + getUsername()
	if r != expectedGreeting {
		t.Errorf("GetGreeting() = %+v, expected %+v", r, expectedGreeting)
	}
}
