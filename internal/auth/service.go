package auth

import (
	"errors"
	"server/internal/user"

	"golang.org/x/crypto/bcrypt"
)

// Service сервич авторизации
type Service struct {
	UserRepository *user.Repository
}

// NewAuthService
func NewAuthService(userRepository *user.Repository) *Service {
	return &Service{UserRepository: userRepository}
}

// Register регистрация нвого пользователя
func (service *Service) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UserRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Password: string(hashedPass),
		Name:     name,
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

// Login логирование пользователя
func (service *Service) Login(email, password string) (string, error) {
	existedUser, err := service.UserRepository.FindByEmail(email)
	if err != nil {
		return "", errors.New(ErrUserNotExists)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))
	//hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New(ErrWrongPassword)
	}
	return existedUser.Email, nil
}
