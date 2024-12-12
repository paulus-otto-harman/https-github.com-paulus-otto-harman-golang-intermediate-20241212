package categoryrepositpry

import (
	"20241212/class/2/domain"
	"20241212/class/2/helper"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepo interface {
	CreateCategory(category *domain.Category) error
	ShowAllCategory(page, limit int) (*[]domain.Category, error)
	DeleteCategory(id int) error
	GetCategoryByID(id int) (*domain.Category, error)
	UpdateCategory(id int, category *domain.Category) error
}

type categoryRepo struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewCategoryRepo(db *gorm.DB, log *zap.Logger) CategoryRepo {
	return &categoryRepo{db, log}
}

func (cr *categoryRepo) ShowAllCategory(page, limit int) (*[]domain.Category, error) {

	category := []domain.Category{}

	result := cr.db.Scopes(helper.Paginate(uint(page), uint(limit))).Find(&category)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("category not found")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (cr *categoryRepo) DeleteCategory(id int) error {

	result := cr.db.Delete(&domain.Category{}, id)

	if result.RowsAffected == 0 {
		return fmt.Errorf("category not found")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *categoryRepo) CreateCategory(category *domain.Category) error {

	err := cr.db.Create(category).Error
	if err != nil {
		return fmt.Errorf("failed to create category: %s", err)
	}

	return nil
}

func (cr *categoryRepo) GetCategoryByID(id int) (*domain.Category, error) {

	category := domain.Category{}
	result := cr.db.Find(&category, id)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("category not found or already deleted")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func (cr *categoryRepo) UpdateCategory(id int, category *domain.Category) error {

	result := cr.db.Model(&category).
		Where("id", id).
		Updates(category)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no record found with shipping_id %d", category.ID)
	}

	return nil
}
