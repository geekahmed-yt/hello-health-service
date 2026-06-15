package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"hello-health-service/internal/app"
)

func TestHealthEndpoint(t *testing.T) {
	handler := app.NewHandler(app.Config{
		Application: "test-app",
		Version:     "1.2.3",
		Commit:      "abc1234",
		Branch:      "feature-x",
	})

	req := httptest.NewRequest(http.MethodGet, "/api/health/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("expected application/json content type, got %q", ct)
	}

	var body map[string]any
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body["status"] != "ok" {
		t.Fatalf("expected status to be ok, got %#v", body["status"])
	}
	if body["application"] != "test-app" {
		t.Fatalf("expected application test-app, got %#v", body["application"])
	}
	if body["version"] != "1.2.3" {
		t.Fatalf("expected version 1.2.3, got %#v", body["version"])
	}
	if body["timestamp"] == "" {
		t.Fatalf("expected timestamp to be present")
	}
}

func TestHelloEndpoint(t *testing.T) {
	handler := app.NewHandler(app.Config{
		Application: "test-app",
		Version:     "2.0.0",
		Commit:      "def9876",
		Branch:      "main",
	})

	req := httptest.NewRequest(http.MethodGet, "/api/hello/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	var body struct {
		Message  string `json:"message"`
		Metadata struct {
			Application string `json:"application"`
			Version     string `json:"version"`
			Commit      string `json:"commit"`
			Branch      string `json:"branch"`
		} `json:"metadata"`
	}
	if err := json.NewDecoder(rr.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if body.Message != "hello world" {
		t.Fatalf("expected message hello world, got %q", body.Message)
	}
	if body.Metadata.Version != "2.0.0" {
		t.Fatalf("expected version 2.0.0, got %q", body.Metadata.Version)
	}
	if body.Metadata.Application != "test-app" {
		t.Fatalf("expected application test-app, got %q", body.Metadata.Application)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	handler := app.NewHandler(app.Config{})

	req := httptest.NewRequest(http.MethodPost, "/api/health/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}
