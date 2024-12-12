package categoryservice

import (
	"20241212/class/2/domain"
	"20241212/class/2/repository"

	"go.uber.org/zap"
)

type CategoryService interface {
	ShowAllCategory(page int) (*[]domain.Category, error)
	CreateCategory(category *domain.Category) error
	DeleteCategory(id int) error
	GetCategoryByID(id int) (*domain.Category, error)
	UpdateCategory(id int, category *domain.Category) error
}

type categoryService struct {
	repo *repository.Repository
	log  *zap.Logger
}

func NewCategoryService(repo *repository.Repository, log *zap.Logger) CategoryService {
	return &categoryService{repo, log}
}

func (cs *categoryService) ShowAllCategory(page int) (*[]domain.Category, error) {

	limit := 20

	categories, err := cs.repo.Category.ShowAllCategory(page, limit)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (cs *categoryService) DeleteCategory(id int) error {

	if err := cs.repo.Category.DeleteCategory(id); err != nil {
		return err
	}

	return nil
}

func (cs *categoryService) CreateCategory(category *domain.Category) error {

	if err := cs.repo.Category.CreateCategory(category); err != nil {
		return err
	}

	return nil
}

func (cs *categoryService) GetCategoryByID(id int) (*domain.Category, error) {

	category, err := cs.repo.Category.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cs *categoryService) UpdateCategory(id int, category *domain.Category) error {

	if err := cs.repo.Category.UpdateCategory(id, category); err != nil {
		return err
	}

	return nil
}
