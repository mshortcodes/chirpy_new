package auth

import "testing"

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("error hashing password: %s", err)
	}

	if CheckPasswordHash(password, hash) != nil {
		t.Fatal("passwords don't match")
	}
}
