package auth

import (
	"errors"
	"net/http"
	"strings"
)

var errBadAuthKey = errors.New("invalid apikey")

func GetApiKey(headers http.Header) (string, error) {
	authH := headers.Get("Authorization")
	if authH == "" {
		return "", errBadAuthKey
	}

	parts := strings.SplitN(authH, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "ApiKey") || strings.TrimSpace(parts[1]) == "" {
		return "", errBadAuthKey
	}

	return parts[1], nil
}
