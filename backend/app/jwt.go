package app

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {
	privateKey, err := generatePrivateKey()
	if err != nil {
		return "", err
	}
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()
	claims["authorized"] = true
	claims["username"] = username
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generatePrivateKey() (ed25519.PrivateKey, error) {
	_, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
