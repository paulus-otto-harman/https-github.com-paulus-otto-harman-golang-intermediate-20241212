package repository

import (
	"20241212/class/2/domain"
	"gorm.io/gorm"
)

type PasswordResetRepository struct {
	db *gorm.DB
}

func NewPasswordResetRepository(db *gorm.DB) *PasswordResetRepository {
	return &PasswordResetRepository{db: db}
}

func (repo PasswordResetRepository) Create(token *domain.PasswordResetToken) error {
	return repo.db.Create(&token).Error
}
