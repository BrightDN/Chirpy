package auth

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

func ComparePasswordHash(password string, hash string) (bool, error) {
	isSame, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, fmt.Errorf("something went wrong: %v", err)
	}
	if !isSame {
		return false, nil
	}
	return true, nil
}
