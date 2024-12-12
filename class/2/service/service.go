package service

import (
	"20241212/class/2/repository"
	categoryservice "20241212/class/2/service/category_service"
	dashboardservice "20241212/class/2/service/dashboard_service"
	productservice "20241212/class/2/service/product_service"

	"go.uber.org/zap"
)

type Service struct {
	Auth          AuthService
	Order         OrderService
	PasswordReset PasswordResetService
	User          UserService
	Category      categoryservice.CategoryService
	Product       productservice.ProductService
	Dashboard     dashboardservice.DashboardService
	Stock         ServiceStock
	Promotion     ServicePromotion
	Banner        ServiceBanner
	Email         EmailService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		Auth:          NewAuthService(repo.Auth),
		Order:         NewOrderService(repo.Order),
		PasswordReset: NewPasswordResetService(repo.PasswordReset),
		User:          NewUserService(repo.User),
		Category:      categoryservice.NewCategoryService(&repo, log),
		Product:       productservice.NewProductService(&repo, log),
		Dashboard:     dashboardservice.NewDashboardService(&repo, log),
		Stock:         NewServiceStock(repo.Stock, log),
		Promotion:     NewServicePromotion(repo.Promotion),
		Banner:        NewServiceBanner(repo.Banner),
		Email:         NewEmailService(),
	}
}
