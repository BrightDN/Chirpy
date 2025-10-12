package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	a := headers.Get("Authorization")
	t := strings.Fields(a)[1]
	if t == "" {
		return t, fmt.Errorf("no bearertoken found")
	}
	return t, nil
}
