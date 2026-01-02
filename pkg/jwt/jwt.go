package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("YogeshNotesApi")

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenstring, err := token.SignedString(SecretKey)
	if err != nil {
		return "", fmt.Errorf("Generate Token Failed: %w", err)
	}

	return tokenstring, nil

}
