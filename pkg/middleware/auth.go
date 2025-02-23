package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/jwt"
)

const (
	ContextEmailKey key = "ContextEmailKey"
)

type key string

func writeUnauthed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func writeAuthContext(r *http.Request, data any) *http.Request {
	ctx := context.WithValue(r.Context(), ContextEmailKey, data)
	return r.WithContext(ctx)
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer ") {
			writeUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		data, isValid := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			writeUnauthed(w)
			return
		}
		req := writeAuthContext(r, data.Email)
		next.ServeHTTP(w, req)
	})
}

func IsAuthedCookie(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			writeUnauthed(w)
			return
		}
		err = cookie.Valid()
		if err != nil {
			writeUnauthed(w)
			return
		}
		data, isValid := jwt.NewJWT(config.Auth.Secret).Parse(cookie.Value)
		if !isValid {
			writeUnauthed(w)
			return
		}
		req := writeAuthContext(r, data.Email)
		next.ServeHTTP(w, req)
	})
}
