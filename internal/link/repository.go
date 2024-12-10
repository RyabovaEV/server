package link

import (
	"server/pkg/db"

	"gorm.io/gorm/clause"
)

// LinkRepository структура с зависимостью от БД
type LinkRepository struct {
	Database *db.Db
}

// NewLinkRepository новый репозиторий
func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

// Create сщдание записи в БД
func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

// GetByHash получение ссылки по хэшу
func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

// Update обновление записи в БД
func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	//Clauses - берет занчение hash из БД если мы его не указали явно
	//есил не использовать clause.Returning{} в ответе json hash будет пустой
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

// Delete удаляем из БД
func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID проверяет запись по id на наличие в БД
func (repo *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	result := repo.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

// Count сколько не удаленных ссылок
func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.Database.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	return count
}

// GetAll получение списка ссылок из БД
func (repo *LinkRepository) GetAll(limit, offset int) []Link {
	var links []Link
	repo.Database.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	return links
}
