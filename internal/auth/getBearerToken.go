package auth

import (
	"errors"
	"net/http"
	"strings"
)

var errBadAuthHeader = errors.New("invalid Authorization header")

func GetBearerToken(headers http.Header) (string, error) {
	authH := headers.Get("Authorization")
	if authH == "" {
		return "", errBadAuthHeader
	}

	parts := strings.SplitN(authH, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") || strings.TrimSpace(parts[1]) == "" {
		return "", errBadAuthHeader
	}

	return parts[1], nil
}
