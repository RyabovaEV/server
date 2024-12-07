package auth

// LoginRecuest структура запроса логина
type LoginRecuest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegistrRequest  структура запроса на регистрацию
type RegistrRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginResponse структура ответа на запрос
type LoginResponse struct {
	Token string `json:"token"`
}

// RegistrResponse структура ответа на запрос
type RegistrResponse struct {
	Token string `json:"token"`
}
