package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

var MySecretKey = []byte("YogeshNotesApi")

func GenerateToken(id int) (string, error) {

	claims := MyClaims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString(MySecretKey)
	if err != nil {
		return "", fmt.Errorf("Generate Token Failed: %w", err)
	}

	return tokenstring, nil

}

func VerifyToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return MySecretKey, nil
	})

	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, nil, errors.New("cannot parse claims")
	}

	return token, claims, nil
}
