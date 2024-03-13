package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"pack-svc/pkg/config"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server interface {
	Start()
}

type appServer struct {
	cfg    config.Config
	lgr    *zap.Logger
	router http.Handler
}

func (s *appServer) Start() {
	server := newHTTPServer(s.cfg.GetHTTPServerConfig(), s.router)

	s.lgr.Sugar().Infof("listening on %s", s.cfg.GetHTTPServerConfig().GetAddress())

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			s.lgr.Sugar().Errorf("failed to start server %+v", err)
		}
	}()

	waitForShutdown(server, s.lgr)
}

func waitForShutdown(server *http.Server, lgr *zap.Logger) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigCh

	defer func() { _ = lgr.Sync() }()

	err := server.Shutdown(context.Background())
	if err != nil {
		lgr.Error(err.Error())
		return
	}

	lgr.Info("server shutdown successful")
}

func newHTTPServer(cfg config.HTTPServerConfig, handler http.Handler) *http.Server {
	return &http.Server{
		Handler:      handler,
		Addr:         cfg.GetAddress(),
		WriteTimeout: time.Second * time.Duration(cfg.GetReadTimeout()),
		ReadTimeout:  time.Second * time.Duration(cfg.GetWriteTimeout()),
	}
}

func NewServer(cfg config.Config, lgr *zap.Logger, router http.Handler) Server {
	return &appServer{
		cfg:    cfg,
		lgr:    lgr,
		router: router,
	}
}
