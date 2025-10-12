package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT_Valid(t *testing.T) {
	userID := uuid.New()
	tok, err := MakeJWT(userID, "secret", time.Hour)
	if err != nil {
		t.Fatalf("make jwt: %v", err)
	}

	got, err := ValidateJWT(tok, "secret")
	if err != nil {
		t.Fatalf("validate: %v", err)
	}
	if got != userID {
		t.Fatalf("want %v, got %v", userID, got)
	}
}
