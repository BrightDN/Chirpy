package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string) (string, error) {
	currentTime := jwt.NewNumericDate(time.Now())
	expireTime := jwt.NewNumericDate(currentTime.Time.Add(time.Hour))
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  currentTime,
		ExpiresAt: expireTime,
		Subject:   userID.String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", fmt.Errorf("something went wrong: %v", err)
	}
	return signed, nil
}
