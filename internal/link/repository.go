package link

import (
	"github.com/user-xat/short-link/internal/models"
	"github.com/user-xat/short-link/pkg/db"

	"gorm.io/gorm/clause"
)

type ILinkRepository interface {
	Create(link *models.Link) (*models.Link, error)
	GetByHash(hash string) (*models.Link, error)
	GetById(id uint) (*models.Link, error)
	GetAll(limit, offset int) []models.Link
	Update(link *models.Link) (*models.Link, error)
	Delete(id uint) error
	Count() int64
}

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *models.Link) (*models.Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*models.Link, error) {
	var link models.Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Update(link *models.Link) (*models.Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&models.Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) GetById(id uint) (*models.Link, error) {
	var link models.Link
	result := repo.Database.DB.First(&link, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) Count() int64 {
	var count int64
	repo.Database.
		Table("links").
		Count(&count)
	return count
}

func (repo *LinkRepository) GetAll(limit, offset int) []models.Link {
	var links []models.Link
	repo.Database.
		Table("links").
		Order("id desc").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	return links
}
