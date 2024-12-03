package req

import "github.com/go-playground/validator/v10"

// IsValid валидация запрашиваемых данных
func IsValid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
