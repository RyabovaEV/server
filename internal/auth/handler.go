// Package auth пакет с хендлером для авторизации
package auth

import (
	"fmt"
	"net/http"
	"server/configs"
	"server/pkg/req"
	"server/pkg/res"
)

// AuthHendlerDeps связь с конфигурацией
type AuthHendlerDeps struct {
	*configs.Config
}

// AuthHendler структура хендлера для авторизации
type AuthHendler struct {
	*configs.Config
}

// NewAuthHendler создание нового хендлера
func NewAuthHendler(router *http.ServeMux, deps AuthHendlerDeps) {
	handler := &AuthHendler{
		Config: deps.Config,
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
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
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
		fmt.Println(*body)
		data := RegistrResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}
