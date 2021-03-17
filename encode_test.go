package atlsnjwt_test

import (
	"fmt"
	"testing"

	atlsnjwt "github.com/AndreyLevchenko/atlassian-jwt"
	"github.com/dgrijalva/jwt-go"
)

const (
	secret    = "top-secret"
	clientKey = "abcd"
)

func Test_Basic(t *testing.T) {
	t.Log("some text")
	ts, encErr := atlsnjwt.Encode("get", "/search?aa=1&bb=2&bb=5&jwt=secret", clientKey, secret, 0)
	if encErr != nil {
		t.Errorf("Encode error: %s", encErr)
	}
	token, err := parseToken(ts, secret)
	if err != nil {
		t.Errorf("Token error: %s", err)
	}

	if !token.Valid {
		t.Errorf("Token isn't valid")
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims["qsh"] != "0bde82f4d2ee508aaf44c7799ac8f66fffde781a00bda6bd5417e979ac571c82" {
		t.Errorf("url hash isn't correct")
	}
	if claims["aud"] != clientKey {
		t.Errorf("'aud' claim isn't correct")
	}
	if claims["iss"] != clientKey {
		t.Errorf("'iss' claim isn't correct")
	}
}

func parseToken(encoded string, secret string) (*jwt.Token, error) {
	return jwt.Parse(encoded, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("token %s\n", token.Header)
		return []byte(secret), nil
	})
}
