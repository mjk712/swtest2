package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler, role string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			fmt.Fprintf(w, "can not find token in header")
			return
		}
		tokenString := r.Header.Get("Token")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("GregorySwtest2"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "error parce token")
			return
		}
		jwtrole := claims["role"].(string)
		if jwtrole != role {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "wait wait wait you are not %s, go away", role)
			return
		}
		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Local().Unix() {
			fmt.Fprintf(w, "token expired")
			return
		}
		id := claims["id"].(float64)

		ctx := context.WithValue(r.Context(), role+"Id", uint64(id))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}