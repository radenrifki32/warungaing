package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"github.com/rifki321/warungku/app"
	"github.com/rifki321/warungku/helper"
)

func TokenAuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
			return
		}
		token, err := app.VerifyToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil {
			helper.Response(w, http.StatusUnauthorized, err.Error(), nil)

			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helper.Response(w, http.StatusUnauthorized, "UNAUTHORIZATION", nil)
			return
		}
		username := claims["username"].(string)
		fmt.Println(username)

		next(w, r, ps)
	}
}
