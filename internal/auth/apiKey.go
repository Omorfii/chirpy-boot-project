package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {

	key_string := headers.Get("Authorization")
	if key_string == "" {
		return "", errors.New("no key found")
	}
	if strings.HasPrefix(key_string, "ApiKey") != true {
		return "", errors.New("unidentified prefix")
	}

	trimed_key := strings.TrimPrefix(key_string, "ApiKey")
	key := strings.TrimSpace(trimed_key)

	return key, nil

}
