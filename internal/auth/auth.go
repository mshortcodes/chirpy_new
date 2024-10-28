package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword uses bcrypt to generate a hashed password from plaintext.
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("couldn't hash password: %s", err)
	}

	return string(hash), nil
}

// CheckPasswordHash checks to compare the password entered by the user with its hashed equivalent.
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return fmt.Errorf("password doesn't match")
	}

	return nil
}

// MakeJWT generates a token and signs it with the passed in secret and SHA-256.
func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   userID.String(),
	})

	signingKey := []byte(tokenSecret)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("couldn't sign token: %s", err)
	}
	return signedToken, nil
}

// ValidateJWT compares two tokens' signatures to validate them and returns the user id.
func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID: %s", err)
	}

	return id, nil
}

// GetBearerToken extracts the token from the Authorization header.
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no auth header included in request")
	}

	splitAuth := strings.Fields(authHeader)
	if splitAuth[0] != "Bearer" || len(splitAuth) < 2 {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

// MakeRefreshToken generates a random 256 bit token encoded in hex.
func MakeRefreshToken() (string, error) {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("couldn't make refresh token: %s", err)
	}

	token := hex.EncodeToString(tokenBytes)
	return token, nil
}

// GetAPIKey extracts the API key from the Authorization header.
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no auth header included in request")
	}

	splitAuth := strings.Fields(authHeader)
	if splitAuth[0] != "ApiKey" || len(splitAuth) < 2 {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}
