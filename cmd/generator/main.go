package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"sync"

	"github.com/samaelkorn/generator/internal/template"
	"github.com/samaelkorn/generator/internal/version"

	"golang.org/x/exp/slog"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	errT := template.DownloadTemplate()

	if errT != nil {
		fmt.Println("test")
		fmt.Println(errT)
	}

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL       string
	httpPort      int
	notifications struct {
		email string
	}
}

type application struct {
	config config
	logger *slog.Logger
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:666", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 666, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.notifications.email, "notifications-email", "", "contact email address for error notifications")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	app := &application{
		config: cfg,
		logger: logger,
	}

	return app.serveHTTP()
}
