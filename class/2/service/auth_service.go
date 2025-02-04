package service

import (
	"20241212/class/2/domain"
	"20241212/class/2/repository"
)

type AuthService interface {
	Login(user domain.User) (string, bool, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(user domain.User) (string, bool, error) {
	return s.repo.Authenticate(user)
}
