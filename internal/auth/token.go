package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {

	token_string := headers.Get("Authorization")
	if token_string == "" {
		return "", errors.New("no token found")
	}
	if strings.HasPrefix(token_string, "Bearer") != true {
		return "", errors.New("unidentified prefix")
	}

	trimed_token := strings.TrimPrefix(token_string, "Bearer")
	token := strings.TrimSpace(trimed_token)

	return token, nil

}
