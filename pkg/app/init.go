package app

import (
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"pack-svc/pkg/config"
	"pack-svc/pkg/http/router"
	"pack-svc/pkg/http/server"
	"pack-svc/pkg/packer"
	"pack-svc/pkg/reporters"
	"pack-svc/pkg/repository"
)

func initHTTPServer(configFile string) {
	cfg := config.NewConfig(configFile)
	logger := initLogger(cfg)
	rt := initRouter(cfg, logger)

	server.NewServer(cfg, logger, rt).Start()
}

func initRouter(cfg config.Config, logger *zap.Logger) http.Handler {
	packRepo := initRepository(cfg)
	packerService := initService(packRepo)

	return router.NewRouter(logger, packerService)
}

func initService(packRepository repository.PackSizesRepository) packer.Service {
	packerService := packer.NewPackerService(packRepository)

	return packerService
}

func initRepository(cfg config.Config) repository.PackSizesRepository {
	return repository.NewPackRepository()
}

func initLogger(cfg config.Config) *zap.Logger {
	return reporters.NewLogger(
		cfg.GetLogConfig().GetLevel(),
		getWriters()...,
	)
}

func getWriters() []io.Writer {
	return []io.Writer{
		os.Stdout,
	}
}
