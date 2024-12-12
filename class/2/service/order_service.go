package service

import (
	"20241212/class/2/domain"
	"20241212/class/2/repository"
)

type OrderService interface {
	Allx(page, limit uint) (int, int, []domain.OrderTotal, error)
	All() ([]domain.OrderTotal, error)
	Update(orderId uint, confirmation domain.OrderConfirmation) error
	Get(orderId uint) (domain.OrderTotal, error)
	Summary() (float32, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) All() ([]domain.OrderTotal, error) {
	return s.repo.All()
}

func (s *orderService) Allx(page, limit uint) (int, int, []domain.OrderTotal, error) {
	return s.repo.Allx(page, limit)
}

func (s *orderService) Update(orderId uint, confirmation domain.OrderConfirmation) error {
	return s.repo.Update(orderId, confirmation)
}

func (s *orderService) Get(orderId uint) (domain.OrderTotal, error) {
	return s.repo.Get(orderId)
}

func (s *orderService) Summary() (float32, error) {
	return s.repo.Summary()
}
