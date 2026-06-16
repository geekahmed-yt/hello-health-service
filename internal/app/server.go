package app

// This file contains the HTTP server implementation for the hello-health-service application.

import (
	"encoding/json"
	"net/http"
	"time"
)

type Config struct {
	Application string
	Version     string
	Commit      string
	Branch      string
}

type healthResponse struct {
	Status      string `json:"status"`
	Application string `json:"application"`
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Branch      string `json:"branch"`
	Timestamp   string `json:"timestamp"`
}

type helloResponse struct {
	Message  string        `json:"message"`
	Metadata helloMetadata `json:"metadata"`
}

type helloMetadata struct {
	Application string `json:"application"`
	Version     string `json:"version"`
	Commit      string `json:"commit"`
	Branch      string `json:"branch"`
}

func NewHandler(cfg Config) http.Handler {
	cfg = normalizeConfig(cfg)

	mux := http.NewServeMux()

	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resp := healthResponse{
			Status:      "ok",
			Application: cfg.Application,
			Version:     cfg.Version,
			Commit:      cfg.Commit,
			Branch:      cfg.Branch,
			Timestamp:   time.Now().UTC().Format(time.RFC3339),
		}
		writeJSON(w, http.StatusOK, resp)
	})

	mux.HandleFunc("/api/hello/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resp := helloResponse{
			Message: "hello world",
			Metadata: helloMetadata{
				Application: cfg.Application,
				Version:     cfg.Version,
				Commit:      cfg.Commit,
				Branch:      cfg.Branch,
			},
		}
		writeJSON(w, http.StatusOK, resp)
	})

	return mux
}

func normalizeConfig(cfg Config) Config {
	if cfg.Application == "" {
		cfg.Application = "hello-health-service"
	}
	if cfg.Version == "" {
		cfg.Version = "dev"
	}
	if cfg.Commit == "" {
		cfg.Commit = "unknown"
	}
	if cfg.Branch == "" {
		cfg.Branch = "local"
	}
	return cfg
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(payload)
}

// This a comment
