package auth

import (
	"net/http"

	"github.com/user-xat/short-link/configs"
	"github.com/user-xat/short-link/pkg/jwt"
	"github.com/user-xat/short-link/pkg/req"
	"github.com/user-xat/short-link/pkg/res"
)

type AuthHandler struct {
	*configs.ApiConfig
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.ApiConfig
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		ApiConfig:   deps.ApiConfig,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		createCookie(w, token)
		res.Json(w, RegisterResponse{
			Token: token,
		}, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := jwt.NewJWT(handler.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		createCookie(w, token)
		res.Json(w, RegisterResponse{
			Token: token,
		}, http.StatusCreated)
	}
}

func createCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    token,
		MaxAge:   86400, // 1 day
		HttpOnly: true,
		Secure:   true,
	})
}
