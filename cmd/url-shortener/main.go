package main

import (
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
	"log"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("Starting server", slog.String("env", cfg.Env))
	log.Debug("Debug is enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to create storage", sl.Err(err))
		os.Exit(1)
	}

	_, err = storage.SaveURL("https://www.google.com", "google")
	if err != nil {
		log.Error("Failed to save url", sl.Err(err))
		os.Exit(1)
	}

	_, err = storage.SaveURL("https://www.google.com", "google")
	if err != nil {
		log.Error("Failed to save url", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
	// TODO: init router: chi, "chi render"

	// TODO: run server:
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
