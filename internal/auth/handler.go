// Package auth пакет с хендлером для авторизации
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
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
		// 1 вариант
		/*reg, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !reg.MatchString(payload.Email) {
			res.Json(w, "Wrong email", 402)
			return
		}*/
		// 2 вариант
		/*match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, payload.Email)
		if !match {
			res.Json(w, "Wrong required", 402)
			return
		}*/
		// 3 вариант через ст библиотеку Go
		_, err = mail.ParseAddress(payload.Email)
		if err != nil {
			res.Json(w, "Wrong required", 402)
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
