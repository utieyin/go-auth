package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	result, err := HashPassword("tony")
	got := []byte(result)
	name := []byte("tony")
	err = bcrypt.CompareHashAndPassword(got, name)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
