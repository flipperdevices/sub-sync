package ghost

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

func generateJwt(apiKey string) (string, error) {
	parts := strings.Split(apiKey, ":")
	if len(parts) != 2 {
		return "", errors.New("wrong api key")
	}
	secret, err := hex.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": fmt.Sprintf("/%s/admin/", VERSION),
		"iat": now.Unix(),
		"exp": now.Add(time.Minute * 2).Unix(),
	})
	token.Header["kid"] = parts[0]

	return token.SignedString(secret)
}
