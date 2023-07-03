package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	secretKey = []byte("RADENMUHAMADRIFKI")
)

func GenerateToken(username string, w http.ResponseWriter) (string, error) {

	fmt.Println(username)

	claims := jwt.MapClaims{
		"exp":        time.Now().Add(10 * time.Hour).Unix(),
		"authorized": true,
		"username":   username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	fmt.Println(tokenString)

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() == jwt.SigningMethodHS256.Alg() {
			return secretKey, nil
		}
		return nil, fmt.Errorf("metode penandatanganan tidak valid: %v", token.Header["alg"])
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
