package main

import (
	todo "Fragaed"
	"Fragaed/internal/config"
	handler "Fragaed/internal/handler"
	"Fragaed/internal/repository"
	"Fragaed/internal/service"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	db, err := repository.NewPostgresDB(*cfg)
	if err != nil {
		log.Info("failed to connect to database", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	if err := server.Run(handlers.InitRoutes(), cfg.HTTPServer.Address, cfg.HTTPServer.Timeout, cfg.HTTPServer.IdleTimeout); err != nil {
		log.Info("failed to start server", err)
	}
	log.Info("starting server")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal, envDev:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
