package repository

import (
	"20241212/class/2/config"
	"20241212/class/2/database"
	categoryrepositpry "20241212/class/2/repository/category_repositpry"
	dashboardrepository "20241212/class/2/repository/dashboard_repository"
	productrepository "20241212/class/2/repository/product_repository"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	Auth          AuthRepository
	Order         OrderRepository
	PasswordReset PasswordResetRepository
	User          UserRepository
	Category      categoryrepositpry.CategoryRepo
	Product       productrepository.ProductRepo
	Dashboard     dashboardrepository.DashboardRepo
	Stock         RepositoryStock
	Promotion     RepositoryPromotion
	Banner        RepositoryBanner
}

func NewRepository(db *gorm.DB, cacher database.Cacher, config config.Config, log *zap.Logger) Repository {
	return Repository{
		Category:      categoryrepositpry.NewCategoryRepo(db, log),
		Product:       productrepository.NewProductRepo(db, log),
		Dashboard:     dashboardrepository.NewDashboardRepo(db, log),
		Auth:          *NewAuthRepository(db, cacher, config.AppSecret),
		Order:         *NewOrderRepository(db),
		PasswordReset: *NewPasswordResetRepository(db),
		User:          *NewUserRepository(db),
		Stock:         NewRepositoryStock(db, log),
		Promotion:     NewRepositoryPromotion(db, log),
		Banner:        *NewRepositoryBanner(db, log),
	}
}
