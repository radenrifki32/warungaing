package helper

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rifki321/warungku/app"
)

func ReadJwt(r *http.Request, w http.ResponseWriter) (string, error) {
	jwtString := r.Header.Get("Authorization")
	if jwtString == "" {
		Response(w, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}
	token, err := app.VerifyToken(strings.TrimPrefix(jwtString, "Bearer "))
	if err != nil {
		Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
	}
	username := claims["username"].(string)
	return username, nil
}
