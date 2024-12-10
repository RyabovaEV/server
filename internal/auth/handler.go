// Package auth обработка HTTP запросов для работы с авторизацией
package auth

import (
	"net/http"
	"server/configs"
	"server/pkg/jwt"
	"server/pkg/req"
	"server/pkg/res"
)

// AuthHendlerDeps связь с конфигурацией
type AuthHendlerDeps struct {
	*configs.Config
	*Service
}

// AuthHendler структура хендлера для авторизации
type AuthHendler struct {
	*configs.Config
	*Service
}

// NewAuthHendler создание нового хендлера
func NewAuthHendler(router *http.ServeMux, deps AuthHendlerDeps) {
	handler := &AuthHendler{
		Config:  deps.Config,
		Service: deps.Service,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

// Login функция вызываемая при логировании
func (handler *AuthHendler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRecuest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.Service.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := LoginResponse{
			Token: token,
		}
		res.Json(w, data, 200)
	}
}

// Register функция вызываемая при регистрации
func (handler *AuthHendler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegistrRequest](&w, r)
		if err != nil {
			return
		}
		email, err := handler.Service.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := RegistrResponse{
			Token: token,
		}
		res.Json(w, data, 200)
	}
}
