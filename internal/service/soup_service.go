package service

import (
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/repository"
)

type SoupService struct {
	repo *repository.SoupRepository
}

func NewSoupService(repo *repository.SoupRepository) *SoupService {
	return &SoupService{repo: repo}
}

func (s *SoupService) Create(soup *model.Soup) error {
	return s.repo.Create(soup)
}

func (s *SoupService) GetByID(id uint) (*model.Soup, error) {
	return s.repo.GetByID(id)
}

func (s *SoupService) Update(soup *model.Soup) error {
	return s.repo.Update(soup)
}

func (s *SoupService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *SoupService) List() ([]model.Soup, error) {
	return s.repo.FindAll()
}
