package user

import "server/pkg/db"

// UserRepository структура с зависимостью от БД
type Repository struct {
	Database *db.Db
}

// NewUserRepository новый репозиторий
func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

// Create сщдание записи в БД
func (repo *Repository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// FindByEmail получение ссылки по хэшу
func (repo *Repository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
