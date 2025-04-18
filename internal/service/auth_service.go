package service

import (
	"errors"
	"uooobarry/soup/internal/auth"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/repository"

	"gorm.io/gorm"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(username, password, email string) (*model.User, error) {
	user := &model.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (*model.User, string, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("user not exist")
		}
		return nil, "", err
	}

	if err := user.ValidatePassword(password); err != nil {
		return nil, "", errors.New("invalid password")
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *AuthService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetUserByID(id)
}
