package main

import (
	"github.com/go-kratos/kratos/v2"
	"lucy/app/engine/internal/conf"
	"lucy/app/engine/internal/data"
	"lucy/app/engine/internal/server"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	database := data.NewMongo(confData)
	dataData, cleanup, err := data.NewData(database, logger)
	if err != nil {
		return nil, nil, err
	}
	cartRepo := data.NewCartRepo(dataData, logger)
	cartUseCase := biz.NewCartUseCase(cartRepo, logger)
	cartService := service.NewCartService(cartUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, logger, tracerProvider, cartService)
	registrar := server.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
