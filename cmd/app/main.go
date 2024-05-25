package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"chi-learn/internals/server"
)

var log *slog.Logger

func init() {
	var level slog.Level
	env := os.Getenv("APP_ENV")

	switch env {
	case "dev":
		level = slog.LevelDebug
	case "prod":
		level = slog.LevelInfo
	default:
		level = slog.LevelWarn
	}

	lh := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})

	log = slog.New(lh)
	slog.SetDefault(log)
}

func main() {
	s := server.New(log)

	port := os.Getenv("PORT")
	if port == "" {
		log.Error("PORT must be set")
		os.Exit(1)
	}

	log.Info(fmt.Sprintf("Listening on port %s", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), s.Router); err != nil {
		log.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
