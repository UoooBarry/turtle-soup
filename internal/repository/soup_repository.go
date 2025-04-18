package repository

import (
	"gorm.io/gorm"
	"uooobarry/soup/internal/model"
)

type SoupRepository struct {
	db *gorm.DB
}

func NewSoupRepository(db *gorm.DB) *SoupRepository {
	return &SoupRepository{db: db}
}

func (r *SoupRepository) Create(soup *model.Soup) error {
	return r.db.Create(soup).Error
}

func (r *SoupRepository) GetByID(id uint) (*model.Soup, error) {
	var soup model.Soup
	if err := r.db.First(&soup, id).Error; err != nil {
		return nil, err
	}
	return &soup, nil
}

func (r *SoupRepository) Update(soup *model.Soup) error {
	return r.db.Save(soup).Error
}

func (r *SoupRepository) Delete(id uint) error {
	return r.db.Delete(&model.Soup{}, id).Error
}

func (r *SoupRepository) FindAll() ([]model.Soup, error) {
	var soups []model.Soup
	if err := r.db.Find(&soups).Error; err != nil {
		return nil, err
	}
	return soups, nil
}
