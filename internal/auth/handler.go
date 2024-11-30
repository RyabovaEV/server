// Package auth пакет с хендлером для авторизации
package auth

import (
	"fmt"
	"net/http"
	"server/configs"
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
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Login")
	}
}

// Register функция вызываемая при регистрации
func (handler *AuthHendler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
