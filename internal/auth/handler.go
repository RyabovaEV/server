// Package auth пакет с хендлером для авторизации
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/configs"
	"server/pcg/res"
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
		var payload LoginRecuest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.Json(w, err.Error(), 402)
			return
		}
		if payload.Email == "" {
			res.Json(w, "Email required", 402)
			return
		}
		if payload.Password == "" {
			res.Json(w, "Password required", 402)
			return
		}
		fmt.Println((payload))
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, 200)
	}
}

// Register функция вызываемая при регистрации
func (handler *AuthHendler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
