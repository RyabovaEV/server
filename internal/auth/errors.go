// Package auth ошибки при авторизации
package auth

const (
	// ErrUserExists пользователеь есть в системе
	ErrUserExists = "user exists"
	// ErrUserNotExists пользователя нет в системе
	ErrUserNotExists = "user not exists"
	// ErrWrongPassword не верный пароль
	ErrWrongPassword = "wrong password"
)
