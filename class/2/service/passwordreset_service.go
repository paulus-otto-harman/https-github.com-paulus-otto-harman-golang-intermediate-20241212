package service

import (
	"20241212/class/2/domain"
	"20241212/class/2/repository"
)

type PasswordResetService interface {
	Create(token *domain.PasswordResetToken) error
}

type passwordResetService struct {
	repo repository.PasswordResetRepository
}

func NewPasswordResetService(repo repository.PasswordResetRepository) PasswordResetService {
	return &passwordResetService{repo: repo}
}

func (s *passwordResetService) Create(token *domain.PasswordResetToken) error {
	return s.repo.Create(token)
}
