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

func IsAuthed(next http.Handler, config *configs.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, ok := isAuthedByHeader(r, config)
		if !ok {
			data, ok = isAuthedByCookie(r, config)
		}
		if !ok {
			writeUnauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func writeUnauthed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func isAuthedByHeader(r *http.Request, config *configs.ApiConfig) (*jwt.JWTData, bool) {
	authedHeader := r.Header.Get("Authorization")
	if !strings.HasPrefix(authedHeader, "Bearer ") {
		return nil, false
	}
	token := strings.TrimPrefix(authedHeader, "Bearer ")
	data, isValid := jwt.NewJWT(config.Auth.Secret).Parse(token)
	if !isValid {
		return nil, false
	}
	return data, true
}

func isAuthedByCookie(r *http.Request, config *configs.ApiConfig) (*jwt.JWTData, bool) {
	cookie, err := r.Cookie("auth")
	if err != nil {
		return nil, false
	}
	err = cookie.Valid()
	if err != nil {
		return nil, false
	}
	data, isValid := jwt.NewJWT(config.Auth.Secret).Parse(cookie.Value)
	if !isValid {
		return nil, false
	}
	return data, true
}
