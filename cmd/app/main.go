package main

import "fmt"
import (
	"fmt"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
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
	fmt.Println("Hello, World!")
}
