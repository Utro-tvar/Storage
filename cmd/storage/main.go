package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/Utro-tvar/Storage/internal/config"
	"github.com/Utro-tvar/Storage/internal/db/sqlite"
	"github.com/Utro-tvar/Storage/internal/lib/logger/sl"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	configPath := flag.String("config", "", "Path to config file")

	flag.Parse()

	cfg := config.MustLoad(*configPath)

	log := setupLogger(cfg.Env)

	log.Info("starting storage", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := sqlite.New(cfg.DBPath)
	if err != nil {
		log.Error("failed to init db", sl.Err(err))
		os.Exit(1)
	}

	_ = db
}

func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log.Fatalf("cannot find logger for environment: %s", env)
	}

	return logger
}
