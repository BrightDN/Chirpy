package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func MakeRefreshToken() (string, error) {
	b := make([]byte, 32)
	n, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("something went wrong: %v", err)
	}
	if n != len(b) {
		return "", fmt.Errorf("the token is not complete: expected a length of %d but got %d", len(b), n)
	}
	hexString := hex.EncodeToString(b)

	return hexString, nil
}
