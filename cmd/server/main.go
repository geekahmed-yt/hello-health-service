package main

import (
	"log"
	"net/http"
	"os"

	"hello-health-service/internal/app"
)

var (
	version = "dev"
	commit  = "unknown"
	branch  = "local"
)

func main() {
	cfg := app.Config{
		Application: envOrDefault("APP_NAME", "hello-health-service"),
		Version:     envOrDefault("APP_VERSION", version),
		Commit:      envOrDefault("APP_COMMIT", commit),
		Branch:      envOrDefault("APP_BRANCH", branch),
	}

	port := envOrDefault("PORT", "8080")
	addr := ":" + port

	log.Printf("starting %s on %s (version=%s branch=%s commit=%s)", cfg.Application, addr, cfg.Version, cfg.Branch, cfg.Commit)
	if err := http.ListenAndServe(addr, app.NewHandler(cfg)); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

// This file is the entry point for the hello-health-service application. It reads configuration from environment variables, sets up an HTTP server, and starts listening for requests. The application responds to health check requests with its version, commit, and branch information.
