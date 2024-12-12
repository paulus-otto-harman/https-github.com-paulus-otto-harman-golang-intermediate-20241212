package infra

import (
	"20241212/class/2/config"
	"20241212/class/2/database"
	"20241212/class/2/handler"
	"20241212/class/2/log"
	"20241212/class/2/middleware"
	"20241212/class/2/repository"
	"20241212/class/2/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cacher     database.Cacher
	Cfg        config.Config
	Ctl        handler.Handler
	Log        *zap.Logger
	Middleware middleware.Middleware
}

func NewServiceContext(migrateDb bool, seedDb bool) (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	appConfig, err := config.LoadConfig(migrateDb, seedDb)
	if err != nil {
		handlerError(err)
	}

	// instance logger
	logger, err := log.InitZapLogger(appConfig)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(appConfig)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(appConfig, 60*60)

	// instance repository
	repo := repository.NewRepository(db, rdb, appConfig, logger)

	// instance service
	service := service.NewService(repo, logger)

	// instance controller
	Ctl := handler.NewHandler(service, logger)

	mw := middleware.NewMiddleware(rdb, appConfig.AppSecret)

	return &ServiceContext{Cacher: rdb, Cfg: appConfig, Ctl: *Ctl, Log: logger, Middleware: mw}, nil
}