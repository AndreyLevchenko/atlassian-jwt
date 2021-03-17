package atlsnjwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	defaultTimeoutSec = 60 * 60
)

func Encode(httpMethod string, url string, clientKey string, sharedSecret string, timeoutSec int) (string, error) {
	if timeoutSec == 0 {
		timeoutSec = defaultTimeoutSec
	}
	secretBytes := []byte(sharedSecret)

	now := int64(time.Now().Unix())

	qsh, err := hashUrl(httpMethod, url)
	if err != nil {
		return "", err
	}

	claims := &jwt.MapClaims{
		"aud": clientKey,
		"exp": now + int64(timeoutSec),
		"iat": now,
		"iss": clientKey,
		"qsh": qsh,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(secretBytes)

	if err != nil {
		return "", err
	}

	return ss, nil
}
